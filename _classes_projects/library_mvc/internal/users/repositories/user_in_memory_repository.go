package repositories

import (
	"errors"
	"sync"
	"time"

	"github.com/felipeolliveira/golang_the_best/_classes_projects/library_mvc/internal/users/models"
)

type UserInMemoryRepository struct {
	users  map[int64]*models.User
	mu     sync.RWMutex
	nextId int64
}

func NewUserInMemoryRepository() *UserInMemoryRepository {
	return &UserInMemoryRepository{
		users:  make(map[int64]*models.User),
		mu:     sync.RWMutex{},
		nextId: 1,
	}
}

func (r *UserInMemoryRepository) CreateUser(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextId
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[user.ID] = user
	r.nextId++

	return nil
}

func (r *UserInMemoryRepository) GetUserById(id int64) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[id]

	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *UserInMemoryRepository) GetUserByEmail(email string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserInMemoryRepository) GetAllUsers() ([]*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	usersList := make([]*models.User, 0, len(r.users))

	for _, user := range r.users {
		usersList = append(usersList, user)
	}

	return usersList, nil
}

func (r *UserInMemoryRepository) UpdateUser(id int64, user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	updateUser, ok := r.users[id]

	if !ok {
		return errors.New("user not found")
	}

	updateUser.Name = user.Name
	updateUser.Email = user.Email
	updateUser.UpdatedAt = time.Now()

	r.users[id] = updateUser

	return nil
}

func (r *UserInMemoryRepository) DeleteUser(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.users, id)

	return nil
}
