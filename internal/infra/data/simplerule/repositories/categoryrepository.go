package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	core.GormRepository[entities.Category, gormmappings.CategoryDB]
}

func NewCategoryRepository(dbe *gorm.DB) domainRepositories.ICategoryRepository {
	return &CategoryRepository{
		core.GormRepository[entities.Category, gormmappings.CategoryDB](*core.NewGormRepository[entities.Category, gormmappings.CategoryDB](dbe)),
	}
} 
