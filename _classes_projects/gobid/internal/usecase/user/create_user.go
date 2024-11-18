package user

import (
	"context"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/gobid/internal/validator"
)

type CreateUserReq struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.UserName), "username", "this field cannot be empty")

	eval.CheckField(validator.NotBlank(req.Bio), "bio", "this field cannot be empty")
	eval.CheckField(validator.MinChars(req.Bio, 10) && validator.MaxChars(req.Bio, 255), "bio", "bio must be between 10 and 255 characters")

	eval.CheckField(validator.MinChars(req.Password, 8), "password", "password must be greater than 8 chars")

	eval.CheckField(validator.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRx), "email", "must be a valid email")

	return eval
}
