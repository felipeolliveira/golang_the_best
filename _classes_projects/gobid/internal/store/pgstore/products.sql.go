// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package pgstore

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
  seller_id,
  product_name,
  description,
  base_price,
  auction_end
)
VALUES ($1, $2, $3, $4, $5)
RETURNING id
`

type CreateProductParams struct {
	SellerID    uuid.UUID      `json:"seller_id"`
	ProductName string         `json:"product_name"`
	Description string         `json:"description"`
	BasePrice   pgtype.Numeric `json:"base_price"`
	AuctionEnd  time.Time      `json:"auction_end"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.SellerID,
		arg.ProductName,
		arg.Description,
		arg.BasePrice,
		arg.AuctionEnd,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}
