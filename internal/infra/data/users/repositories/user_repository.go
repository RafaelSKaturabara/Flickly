package repositories

import (
	"errors"
	"flickly/internal/domain/users/entities"
)

type UserRepository struct {
	Users []entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	for _, existingUser := range r.Users {
		if user.Email == existingUser.Email {
			return errors.New("user already exists")
		}
	}
	r.Users = append(r.Users, *user)
	return nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}
