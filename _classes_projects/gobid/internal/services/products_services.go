package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

var ProductServiceErr = struct {
	CouldNotBeCreated error
}{
	errors.New("product could not be created"),
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool,
		pgstore.New(pool),
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, sellerId uuid.UUID, productName, description string, baseprice float64, auctionEnd time.Time) (uuid.UUID, error) {
	var basePriceNum pgtype.Numeric

	if err := basePriceNum.Scan(baseprice); err != nil {
		slog.Error("products_services", "CreateProduct", err)
		return uuid.UUID{}, ProductServiceErr.CouldNotBeCreated
	}

	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		SellerID:    sellerId,
		ProductName: productName,
		Description: description,
		BasePrice:   basePriceNum,
		AuctionEnd:  auctionEnd,
	})
	if err != nil {
		slog.Error("products_services", "CreateProduct", err)
		return uuid.UUID{}, ProductServiceErr.CouldNotBeCreated
	}

	return id, nil
}
