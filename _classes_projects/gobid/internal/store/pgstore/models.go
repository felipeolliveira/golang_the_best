// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgstore

import (
	"time"

	"github.com/google/uuid"
)

type Bid struct {
	ID               uuid.UUID `json:"id"`
	ProductID        uuid.UUID `json:"product_id"`
	BidderID         uuid.UUID `json:"bidder_id"`
	BidAmountInCents int32     `json:"bid_amount_in_cents"`
	CreatedAt        time.Time `json:"created_at"`
}

type Product struct {
	ID               uuid.UUID `json:"id"`
	SellerID         uuid.UUID `json:"seller_id"`
	ProductName      string    `json:"product_name"`
	Description      string    `json:"description"`
	BasePriceInCents int32     `json:"base_price_in_cents"`
	AuctionEnd       time.Time `json:"auction_end"`
	IsSold           bool      `json:"is_sold"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Session struct {
	Token  string    `json:"token"`
	Data   []byte    `json:"data"`
	Expiry time.Time `json:"expiry"`
}

type SubscribeToken struct {
	UserID uuid.UUID `json:"user_id"`
	Token  uuid.UUID `json:"token"`
	Exp    time.Time `json:"exp"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"password_hash"`
	Bio          string    `json:"bio"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
