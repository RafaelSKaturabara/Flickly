package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
	"github.com/rkaturabara/flickly/internal/infra/data/core"
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