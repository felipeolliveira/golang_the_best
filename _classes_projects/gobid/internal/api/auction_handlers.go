package api

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/jsonutils"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) handleSubscribeUserToAuction(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "productId")

	productId, err := uuid.Parse(rawProductId)
	if err != nil {
		jsonutils.EncodeJsonWithError(w, r, http.StatusBadRequest, "invalid product id - must be a valid UUID")
		return
	}

	_, err = api.ProductService.GetProductById(r.Context(), productId)
	if err != nil {
		if errors.Is(err, services.ProductServiceErr.NotFoundProduct) {
			jsonutils.EncodeJson(w, r, http.StatusNotFound, err.Error())
			return
		}

		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	api.AuctionLobby.Lock()
	room, ok := api.AuctionLobby.Rooms[productId]
	api.AuctionLobby.Unlock()

	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"message": "the auction has ended",
		})
	}

	userId, ok := api.Session.Get(r.Context(), SessionUserKey).(uuid.UUID)
	if !ok {
		slog.Error("auction_handlers", "handleGetTokenToSubscribeToAuction", "user not found in session")
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, "internal server error")
		return
	}

	conn, err := api.WsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("auction_handlers", "update connection err", err)
		jsonutils.EncodeJsonWithError(w, r, http.StatusBadRequest, "You must use a websocket connection")
		return
	}

	client := services.NewClient(room, conn, userId)

	room.Register <- client

	go client.ReadEventLoop()
	go client.WriteEventLoop()
}
