package application

import "github.com/google/uuid"

type contextKey string

const UserContextKey contextKey = "user"

type UserAuth struct {
	ID    uuid.UUID
	Email string
	Name  string
	Roles []string
}

func NewUserAuth(id, email, name string, roles []string) UserAuth {
	return UserAuth{
		ID:    uuid.MustParse(id),
		Email: email,
		Name:  name,
		Roles: roles,
	}
}