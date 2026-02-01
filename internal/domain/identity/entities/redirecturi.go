package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type RedirectURI struct {
	core.BaseEntity
	Value string
	ClientID uuid.UUID
	Client   *Client
}

func NewRedirectURI(value string) RedirectURI {
	return RedirectURI{
		BaseEntity: core.NewBaseEntity(),
		Value:      value,
	}
}

func (r *RedirectURI) IsValid() bool {
	return r.ClientID != uuid.Nil
}
