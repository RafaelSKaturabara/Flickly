package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type SubcategoryRepository struct {
	core.GormRepository[entities.Subcategory, gormmappings.SubcategoryDB]
}

func NewSubcategoryRepository(dbe *gorm.DB) domainRepositories.ISubcategoryRepository {
	return &SubcategoryRepository{
		core.GormRepository[entities.Subcategory, gormmappings.SubcategoryDB](*core.NewGormRepository[entities.Subcategory, gormmappings.SubcategoryDB](dbe)),
	}
} 
