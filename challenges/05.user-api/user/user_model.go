package user

import (
	"time"
)

type UserModel struct {
	FirstName string
	LastName  string
	biography string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}

func (u *UserModel) IsDeleted() bool {
	return !u.deletedAt.IsZero()
}

func (u *UserModel) Delete() {
	u.deletedAt = time.Now()
}

func (u *UserModel) Update(firstName string, lastName string, bio string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.biography = bio
	u.updatedAt = time.Now()
}

func NewUser(firstName string, lastName string, bio string) UserModel {
	return UserModel{
		FirstName: firstName,
		LastName:  lastName,
		biography: bio,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
