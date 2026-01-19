package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
	"github.com/rkaturabara/flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	core.GormRepository[entities.Category]
}

func NewCategoryRepository(dbe *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		core.GormRepository[entities.Category](*core.NewGormRepository[entities.Category](dbe)),
	}
} 