package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type Subcategory struct {
	core.BaseEntity
	CategoryID  int64  `json:"category_id"`
	Name        string `json:"name"`        // "Shelter", "Dining Out", etc.
	Description string `json:"description"` // Texto explicativo das suas imagens
}

func (u *Subcategory) IsValid() bool {
	return true
}
