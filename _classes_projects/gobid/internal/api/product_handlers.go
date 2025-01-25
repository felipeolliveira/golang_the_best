package api

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/jsonutils"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/services"
	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/usecase/product"
	"github.com/google/uuid"
)

func (api *Api) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[product.CreateProductReq](r)
	if len(problems) > 0 {
		jsonutils.EncodeJsonWithError(w, r, http.StatusUnprocessableEntity, problems)
		return
	}
	if err != nil {
		jsonutils.EncodeJsonWithError(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userId, ok := api.Session.Get(r.Context(), SessionUserKey).(uuid.UUID)
	if !ok {
		slog.Error("product_handlers", "handleCreatedProduct", "error on get authenticated user id")
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, "internal server error, try again later")
		return
	}

	productId, err := api.ProductService.CreateProduct(r.Context(), userId, data.ProductName, data.Description, data.BasePrise, data.AuctionEnd)
	if err != nil {
		jsonutils.EncodeJsonWithError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	ctx, _ := context.WithDeadline(context.Background(), data.AuctionEnd)
	auctionRoom := services.NewActionRoom(ctx, productId, api.BidsService)

	go auctionRoom.Run()

	api.AuctionLobby.Lock()
	api.AuctionLobby.Rooms[productId] = auctionRoom
	api.AuctionLobby.Unlock()

	jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"productId": productId,
		"auction":   "Room started successfully",
	})
}
