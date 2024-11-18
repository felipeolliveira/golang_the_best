package user

import (
	"context"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/validator"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.Matches(req.Email, validator.EmailRx), "email", "must be a valid email")
	eval.CheckField(validator.NotBlank(req.Password), "password", "this field cannot be empty")

	return eval
}
