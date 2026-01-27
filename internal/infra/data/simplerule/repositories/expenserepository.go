package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type ExpenseRepository struct {
	core.GormRepository[entities.Expense, gormmappings.ExpenseDB]
}

func NewExpenseRepository(dbe *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{
		core.GormRepository[entities.Expense, gormmappings.ExpenseDB](*core.NewGormRepository[entities.Expense, gormmappings.ExpenseDB](dbe)),
	}
} 
