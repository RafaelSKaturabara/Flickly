package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type IncomeRepository struct {
	core.GormRepository[entities.Income]
}

func NewIncomeRepository(dbe *gorm.DB) *IncomeRepository {
	return &IncomeRepository{
		core.GormRepository[entities.Income](*core.NewGormRepository[entities.Income](dbe)),
	}
} 
