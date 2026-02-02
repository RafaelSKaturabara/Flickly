package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type Role struct {
	core.BaseEntity
	Value    string
	ClientID *uuid.UUID // null indica que é default para todos os clientes
	Client   *Client    // null indica que é default para todos os clientes
}

func NewRole(value string) Role {
	return Role{
		BaseEntity: core.NewBaseEntity(),
		Value:      value,
	}
}

func (r *Role) IsValid() bool {
	return true
}

const (
	RoleAdmin string = "admin"
	RoleUser  string = "user"
)
