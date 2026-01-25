package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	core.GormRepository[entities.Category, gormmappings.CategoryDB]
}

func NewCategoryRepository(dbe *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		core.GormRepository[entities.Category, gormmappings.CategoryDB](*core.NewGormRepository[entities.Category, gormmappings.CategoryDB](dbe)),
	}
} 
