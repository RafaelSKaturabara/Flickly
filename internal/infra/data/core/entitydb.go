package core

import (
	"time"

	"github.com/google/uuid"
)

type EntityDB struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdateAt  *time.Time `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"index"`
}