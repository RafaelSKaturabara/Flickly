package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type UserAccess struct {
	core.BaseEntity
	UserID   uuid.UUID
	User     *User
	ClientID uuid.UUID
	Client   *Client
	Roles    []Role
}

func NewUserAccess(user *User, client *Client, roles []Role) UserAccess {
	var userID, ClientID uuid.UUID

	if user != nil {
		userID = user.ID
	}

	if client != nil {
		ClientID = client.ID
	}

	return UserAccess{
		BaseEntity: core.NewBaseEntity(),
		UserID:     userID,
		User:       user,
		ClientID:   ClientID,
		Client:     client,
		Roles:      roles,
	}
}

func (ua *UserAccess) IsValid() bool {
	return ua.UserID != uuid.Nil && ua.ClientID != uuid.Nil
}
