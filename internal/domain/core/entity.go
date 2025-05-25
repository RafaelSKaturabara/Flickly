package core

import (
	"time"

	"github.com/google/uuid"
)

type Entity interface {
	GetID() uuid.UUID
	GetCreatedAt() time.Time
	GetUpdateAt() *time.Time
	GetDeletedAt() *time.Time
}

type BaseEntity struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdateAt  *time.Time `json:"updateAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}

func (e *BaseEntity) GetID() uuid.UUID {
	return e.ID
}

func (e *BaseEntity) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *BaseEntity) GetUpdateAt() *time.Time {
	return e.UpdateAt
}

func (e *BaseEntity) GetDeletedAt() *time.Time {
	return e.DeletedAt
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
