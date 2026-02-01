package entities

import (
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type AccessGrant struct {
	core.BaseEntity
	UserID    uuid.UUID
	User      *User
	ClientID  uuid.UUID
	Client    *Client
	TokenHash string // O segredo (SHA256)
	ExpiresAt time.Time
	Revoked   bool // Se foi revogado. Ex: Logout
	// Metadados para o usuário saber de onde logou
	UserAgent string // Navegador e SO
	IPAddress string // IP do dispositivo
}

func NewAccessGrant(user *User, client *Client, tokenHash string, expiresAt time.Time, revoked bool, userAgent string, ipAddress string) AccessGrant {
	return AccessGrant{
		BaseEntity: core.NewBaseEntity(),
		UserID:     user.ID,
		User:       user,
		ClientID:   client.ID,
		Client:     client,
		TokenHash:  tokenHash,
		ExpiresAt:  expiresAt,
		Revoked:    revoked,
	}
}

func (ag *AccessGrant) IsValid() bool {
	return ag.UserID != uuid.Nil && ag.ClientID != uuid.Nil && !ag.Revoked && ag.ExpiresAt.After(time.Now())
}
