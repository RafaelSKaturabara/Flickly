package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
	"github.com/rkaturabara/flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	core.GormRepository[entities.Expense]
}

func NewExpenseRepository(dbe *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{
		core.GormRepository[entities.Expense](*core.NewGormRepository[entities.Expense](dbe)),
	}
} 