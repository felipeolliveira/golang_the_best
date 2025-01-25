package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

var ProductServiceErr = struct {
	CouldNotBeCreated error
	NotFoundProduct   error
}{
	errors.New("product could not be created"),
	errors.New("not found product by id"),
}

func NewProductService(pool *pgxpool.Pool) ProductService {
	return ProductService{
		pool,
		pgstore.New(pool),
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, sellerId uuid.UUID, productName, description string, basePriceInCents int, auctionEnd time.Time) (uuid.UUID, error) {
	id, err := ps.queries.CreateProduct(ctx, pgstore.CreateProductParams{
		SellerID:         sellerId,
		ProductName:      productName,
		Description:      description,
		BasePriceInCents: int32(basePriceInCents),
		AuctionEnd:       auctionEnd,
	})
	if err != nil {
		slog.Error("products_services", "CreateProduct", err)
		return uuid.UUID{}, ProductServiceErr.CouldNotBeCreated
	}

	return id, nil
}

func (ps *ProductService) GetProductById(ctx context.Context, productId uuid.UUID) (pgstore.Product, error) {
	product, err := ps.queries.GetProductById(ctx, productId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Product{}, ProductServiceErr.NotFoundProduct
		}

		slog.Error("products_services", "GetProductById", err)
		return pgstore.Product{}, ProductServiceErr.CouldNotBeCreated
	}

	return product, nil
}
