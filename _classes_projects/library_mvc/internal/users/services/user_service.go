package services

import (
	"errors"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/models"
)

type UserService struct {
	userRepo models.UserRepository
}

func NewUserService(userRepo models.UserRepository) models.UserService {
	return &UserService{userRepo: userRepo}
}

func (s UserService) CreateUser(user *models.User) error {
	foundUser, _ := s.userRepo.GetUserByEmail(user.Email)

	if foundUser != nil {
		return errors.New("user with this email already exists")
	}

	return s.userRepo.CreateUser(user)
}

func (s UserService) GetUser(id int64) (*models.User, error) {
	return s.userRepo.GetUserById(id)
}

func (s UserService) GetAllUsers() ([]*models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s UserService) UpdateUser(id int64, user *models.User) error {
	return s.userRepo.UpdateUser(id, user)
}

func (s UserService) DeleteUser(id int64) error {
	return s.userRepo.DeleteUser(id)
}
