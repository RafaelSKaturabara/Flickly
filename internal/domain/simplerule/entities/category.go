package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type Category struct {
	core.BaseEntity
	Name string `json:"name"` // "Needs", "Wants", "Savings"
}

func (u *Category) IsValid() bool {
	return true
}
