package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
)

type ISubcategoryRepository interface {
	core.Repository[entities.Subcategory]
}
