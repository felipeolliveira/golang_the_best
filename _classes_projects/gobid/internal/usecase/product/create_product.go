package product

import (
	"context"
	"time"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/validator"
)

type CreateProductReq struct {
	AuctionEnd  time.Time `json:"auctionEnd"`
	ProductName string    `json:"productName"`
	Description string    `json:"description"`
	BasePrise   int       `json:"basePriceInCents"`
}

const (
	minAuctionDuration  = 2 * time.Hour
	maxBasePriceInCents = 100000000 // 1 million of money / 100 million of cents
)

func (req CreateProductReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.ProductName), "productName", "this field cannot be blank")

	eval.CheckField(validator.NotBlank(req.Description), "description", "this field cannot be blank")
	eval.CheckField(validator.MinChars(req.Description, 10) && validator.MaxChars(req.Description, 255), "description", "this field must have between 10 and 255 characters")

	eval.CheckField(req.BasePrise <= maxBasePriceInCents, "basePriceInCents", "this field must be less than a million")
	eval.CheckField(req.BasePrise > 0, "basePriceInCents", "this field must be greater than 0")

	eval.CheckField(time.Until(req.AuctionEnd) >= minAuctionDuration, "auctionEnd", "auction must last at least 2 hours")

	return eval
}
