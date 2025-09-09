package models

type UserRepository interface {
	CreateUser(user *User) error
	GetUserById(id int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(id int64, user *User) error
	DeleteUser(id int64) error
}
