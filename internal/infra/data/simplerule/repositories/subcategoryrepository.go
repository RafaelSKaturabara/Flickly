package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
	"github.com/rkaturabara/flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type SubcategoryRepository struct {
	core.GormRepository[entities.Subcategory]
}

func NewSubcategoryRepository(dbe *gorm.DB) *SubcategoryRepository {
	return &SubcategoryRepository{
		core.GormRepository[entities.Subcategory](*core.NewGormRepository[entities.Subcategory](dbe)),
	}
} 