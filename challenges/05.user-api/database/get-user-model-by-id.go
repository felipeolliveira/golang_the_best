package database

import (
	"github.com/felipeolliveira/user_api_golang_challenge/user"
	"github.com/google/uuid"
)

type userModelWithId struct {
	Id uuid.UUID
	user.UserModel
}

func getUserModelById(id string) (userModelWithId, error) {
	parsedUuid, err := uuid.Parse(id)
	if err != nil {
		return userModelWithId{}, ErrParsingUuid
	}

	foundedUser, ok := db[parsedUuid]

	if !ok {
		return userModelWithId{}, ErrNotFoundUserFromId
	}

	return userModelWithId{
		Id:        parsedUuid,
		UserModel: foundedUser,
	}, nil
}
