package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
)

type Subcategory struct {
	core.BaseEntity
	Name             string
	Description      string
	CategoryID       uuid.UUID
	Category         Category
	UserSimpleRuleID *uuid.UUID
}

func NewSubcategory(name string, description string, categoryID uuid.UUID) Subcategory {
	return Subcategory{
		BaseEntity:  core.NewBaseEntity(),
		Name:        name,
		Description: description,
		CategoryID:  categoryID,
	}
}

func (u *Subcategory) IsValid() bool {
	return true
}
