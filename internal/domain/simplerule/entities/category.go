package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type Category struct {
	core.BaseEntity
	Name string
}

func NewCategory(name string) Category {
	return Category{
		BaseEntity: core.NewBaseEntity(),
		Name:       name,
	}
}

func (u *Category) IsValid() bool {
	return true
}
