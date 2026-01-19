package entities

import "github.com/rkaturabara/flickly/internal/domain/core"

type Subcategory struct {
	core.BaseEntity
	CategoryID  int64  `json:"category_id"`
	Name        string `json:"name"`        // "Shelter", "Dining Out", etc.
	Description string `json:"description"` // Texto explicativo das suas imagens
}

func (u *Subcategory) IsValid() bool {
	return true
}