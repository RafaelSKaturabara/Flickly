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
	IsValid() bool
	GetErrors() *map[string]string
}

type BaseEntity struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdateAt  *time.Time
	DeletedAt *time.Time
	Errors    *map[string]string `gorm:"-"`
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

func (e *BaseEntity) GetErrors() *map[string]string {
	return e.Errors
}

func NewBaseEntity() BaseEntity {
	return BaseEntity{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
