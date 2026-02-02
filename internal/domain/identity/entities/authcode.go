package entities

import (
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type AuthCode struct {
	core.BaseEntity
	Code          string   
	ClientID      uuid.UUID 
	Client        *Client
	UserID        *uuid.UUID
	User          *User
	CodeChallenge string
	ExpiresAt     time.Time
}

func NewAuthCode(code string, clientID uuid.UUID, userID *uuid.UUID, codeChallenge string, expiresAt time.Time) *AuthCode {
	return &AuthCode{
		BaseEntity: core.NewBaseEntity(),
		Code:          code,
		ClientID:      clientID,
		UserID:        userID,
		CodeChallenge: codeChallenge,
		ExpiresAt:     expiresAt,
	}
}

func (a *AuthCode) IsValid() bool {
	return true
}
