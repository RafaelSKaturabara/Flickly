package repositories

import (
	"flickly/internal/domain/users/entities"
)

type IUserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}
