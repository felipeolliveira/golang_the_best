package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/store/pgstore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BidsService struct {
	pool    *pgxpool.Pool
	queries *pgstore.Queries
}

var BidsServiceErr = struct {
	NotFoundProductById error
	CouldNotPlaceBid    error
	BidIsTooLow         error
}{
	errors.New("product not found"),
	errors.New("could not place bid"),
	errors.New("the bid values is too low"),
}

func NewBidsService(pool *pgxpool.Pool) BidsService {
	return BidsService{
		pool:    pool,
		queries: pgstore.New(pool),
	}
}

func (bs *BidsService) Placebid(ctx context.Context, productId, bidderId uuid.UUID, amountInCents int32) (pgstore.Bid, error) {
	product, err := bs.queries.GetProductById(ctx, productId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return pgstore.Bid{}, BidsServiceErr.NotFoundProductById
		}

		slog.Error("bids_service", "Placebid::GetProductById", err)
		return pgstore.Bid{}, BidsServiceErr.CouldNotPlaceBid
	}

	highestBid, err := bs.queries.GetHighestBidByProductId(ctx, product.ID)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			slog.Error("bids_service", "Placebid::GetHighestBidByProductId", err)
			return pgstore.Bid{}, BidsServiceErr.CouldNotPlaceBid
		}
	}

	if product.BasePriceInCents >= amountInCents || highestBid.BidAmountInCents >= amountInCents {
		return pgstore.Bid{}, BidsServiceErr.BidIsTooLow
	}

	highestBid, err = bs.queries.CreateBid(ctx, pgstore.CreateBidParams{
		ProductID:        productId,
		BidderID:         bidderId,
		BidAmountInCents: amountInCents,
	})
	if err != nil {
		slog.Error("bids_service", "Placebid::CreateBid", err)
		return pgstore.Bid{}, BidsServiceErr.CouldNotPlaceBid
	}

	return highestBid, nil
}
