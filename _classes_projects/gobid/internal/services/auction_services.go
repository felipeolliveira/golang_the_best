package services

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	Send   chan Message
	Room   *AuctionRoom
	UserId uuid.UUID
}

type MessageKind int

const (
	PlaceBid MessageKind = iota

	// Success
	SuccessFullyPlaceBid

	// Errors
	FailedToPlaceBid
	InvalidJSON

	// Info
	NewBidPlaced
	AuctionFinished
)

const (
	maxMessageSize = 512
	readDeadLine   = 60 * time.Second
	pingPeriod     = (readDeadLine * 9) / 10 // 90% of readDeadLine
	writeWait      = 10 * time.Second
)

type Message struct {
	Message       string      `json:"message,omitempty"`
	Kind          MessageKind `json:"kind"`
	UserId        uuid.UUID   `json:"userId,omitempty"`
	AmountInCents int32       `json:"amountInCents,omitempty"`
}

// Contexto da AuctionRoom servirá para parar a execução da goroutine que escuta por novas mensagens
// Channel de Broadcast servirá para receber e enviar mensagens para os clients conectados
// Register e Unregister servirão para adicionar e remover clients, da AuctionRoom, uma vez que:
// - Clients é uma struct e pode sofrer race conditions quando os clients são adicionados ou removidos
type AuctionRoom struct {
	Context     context.Context
	BidsService BidsService
	Broadcast   chan Message
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[uuid.UUID]*Client
	Id          uuid.UUID
}

func (r *AuctionRoom) registerClient(client *Client) {
	slog.Info("New user connected", "Client", client)
	r.Clients[client.UserId] = client
}

func (r *AuctionRoom) unregisterClient(client *Client) {
	slog.Info("New user disconnected", "Client", client)
	delete(r.Clients, client.UserId)
}

func (r *AuctionRoom) broadcastMessage(message Message) {
	slog.Info("Broadcasting message", "RoomId", r.Id, "userId", message.UserId, "Message", message.Message, "Kind", message.Kind)

	switch message.Kind {
	case PlaceBid:
		bid, err := r.BidsService.Placebid(r.Context, r.Id, message.UserId, message.AmountInCents)
		if err != nil {
			if errors.Is(err, BidsServiceErr.BidIsTooLow) {
				if client, ok := r.Clients[message.UserId]; ok {
					client.Send <- Message{Kind: FailedToPlaceBid, Message: err.Error(), UserId: message.UserId}
				}
			}

			slog.Error("auction_services", "broadcastMessage", err)
			return
		}

		if client, ok := r.Clients[message.UserId]; ok {
			client.Send <- Message{Kind: SuccessFullyPlaceBid, Message: "Bid placed successfully", UserId: message.UserId}
		}

		for id, client := range r.Clients {
			if id == message.UserId {
				continue
			}

			newBidMessage := Message{Kind: NewBidPlaced, Message: "New bid placed", AmountInCents: bid.BidAmountInCents, UserId: message.UserId}
			client.Send <- newBidMessage
		}
	case InvalidJSON:
		client, ok := r.Clients[message.UserId]
		if !ok {
			slog.Info("Client not found", "UserId", message.UserId)
			return
		}
		client.Send <- message
	}
}

func (r *AuctionRoom) Run() {
	slog.Info("Starting auction room", "AuctionRoomId", r.Id)

	defer func() {
		close(r.Broadcast)
		close(r.Register)
		close(r.Unregister)
	}()

	for {
		select {
		case client := <-r.Register:
			r.registerClient(client)

		case client := <-r.Unregister:
			r.unregisterClient(client)

		case message := <-r.Broadcast:
			r.broadcastMessage(message)

		case <-r.Context.Done():
			slog.Info("Auction room closed", "AuctionRoomId", r.Id)
			for _, client := range r.Clients {
				client.Send <- Message{Kind: AuctionFinished, Message: "Auction room has been finished"}
			}
			return
		}
	}
}

type AuctionLobby struct {
	Rooms map[uuid.UUID]*AuctionRoom
	sync.Mutex
}

func NewClient(room *AuctionRoom, conn *websocket.Conn, userId uuid.UUID) *Client {
	return &Client{
		Room:   room,
		Conn:   conn,
		Send:   make(chan Message, 512),
		UserId: userId,
	}
}

func NewActionRoom(ctx context.Context, id uuid.UUID, bidsService BidsService) *AuctionRoom {
	return &AuctionRoom{
		Id:          id,
		Broadcast:   make(chan Message),
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[uuid.UUID]*Client),
		Context:     ctx,
		BidsService: bidsService,
	}
}

func NewAuctionLobby() AuctionLobby {
	return AuctionLobby{
		Rooms: make(map[uuid.UUID]*AuctionRoom),
	}
}

func (c *Client) ReadEventLoop() {
	defer func() {
		c.Room.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(readDeadLine))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(readDeadLine))
		return nil
	})

	for {
		m := Message{
			UserId: c.UserId,
		}
		err := c.Conn.ReadJSON(&m)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				slog.Error("auction_services", "ReadEventLoop::websocket", err)
				return
			}

			slog.Error("auction_services", "ReadEventLoop::err", err)
			c.Room.Broadcast <- Message{
				Kind:    InvalidJSON,
				Message: "this message should be a valid JSON",
				UserId:  m.UserId,
			}

			continue
		}

		c.Room.Broadcast <- m
	}
}

func (c *Client) WriteEventLoop() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteJSON(Message{
					Kind:    websocket.CloseMessage,
					Message: "Client is disconnected",
				})
				return
			}
			if message.Kind == AuctionFinished {
				close(c.Send)
				return
			}
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

			err := c.Conn.WriteJSON(message)
			if err != nil {
				c.Room.Unregister <- c
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				slog.Error("auction_services", "WriteEventLoop", err)
				return
			}
		}
	}
}
