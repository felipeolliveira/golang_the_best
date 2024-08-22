package user

import (
	"time"

	"github.com/google/uuid"
)

type UserView struct {
	Id        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func NewUserView(id uuid.UUID, userModel UserModel) UserView {
	return UserView{
		Id:        id,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Biography: userModel.biography,
		CreatedAt: userModel.createdAt.UTC().Format(time.RFC3339),
		UpdatedAt: userModel.updatedAt.UTC().Format(time.RFC3339),
	}
}
