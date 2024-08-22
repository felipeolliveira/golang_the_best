package database

import (
	"errors"
	"sync"

	"github.com/felipeolliveira/user_api_golang_challenge/user"
	"github.com/google/uuid"
)

var mu sync.Mutex
var db = make(map[uuid.UUID]user.UserModel)

var (
	ErrNotFoundUserFromId    = errors.New("the user with the specified ID does not exist")
	ErrCouldNotRetrieveUsers = errors.New("the users information could not be retrieved")
	ErrCouldNotSaveUserInDB  = errors.New("there was an error while saving the user to the database")
	ErrCouldNotBeModifyUser  = errors.New("the user information could not be modified")
	ErrParsingUuid           = errors.New("error when parsing uuid")
)
