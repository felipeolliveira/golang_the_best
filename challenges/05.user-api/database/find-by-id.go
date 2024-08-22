package database

import "github.com/felipeolliveira/user_api_golang_challenge/user"

func FindById(id string) (user.UserView, error) {
	userWithId, err := getUserModelById(id)
	if err != nil {
		return user.UserView{}, err
	}

	return user.NewUserView(userWithId.Id, userWithId.UserModel), nil
}
