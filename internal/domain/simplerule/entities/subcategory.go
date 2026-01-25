package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type Subcategory struct {
	core.BaseEntity
	Name        string
	Description string
	CategoryID  uuid.UUID
	Category    Category
}

func (u *Subcategory) IsValid() bool {
	return true
}
