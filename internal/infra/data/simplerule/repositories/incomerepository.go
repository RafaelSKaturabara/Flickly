package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type IncomeRepository struct {
	core.GormRepository[entities.Income, gormmappings.IncomeDB]
}

func NewIncomeRepository(dbe *gorm.DB) domainRepositories.IIncomeRepository {
	return &IncomeRepository{
		core.GormRepository[entities.Income, gormmappings.IncomeDB](*core.NewGormRepository[entities.Income, gormmappings.IncomeDB](dbe)),
	}
} 
