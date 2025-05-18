package entities

import (
	"flickly/internal/domain/core"
)

type User struct {
	core.Entity
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name string, email string) *User {
	return &User{
		Entity: core.NewEntity(),
		Name:   name,
		Email:  email,
	}
}
