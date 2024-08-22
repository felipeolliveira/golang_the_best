package database

import (
	"github.com/felipeolliveira/user_api_golang_challenge/user"
)

func FindAll() ([]user.UserView, error) {
	parsedUsers := make([]user.UserView, 0, len(db))

	for id, userInDb := range db {
		if !userInDb.IsDeleted() {
			parsedUsers = append(parsedUsers, user.NewUserView(id, userInDb))
		}
	}

	return parsedUsers, nil
}
