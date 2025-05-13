package core

import (
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID           uuid.UUID  `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastUpdateAt *time.Time `json:"lastUpdateAt,omitempty"`
	DeletedAt    *time.Time `json:"deletedAt,omitempty"`
}

func NewEntity() Entity {
	return Entity{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}
}
