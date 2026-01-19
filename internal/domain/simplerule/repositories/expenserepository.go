package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
)

type IExpenseRepository interface {
	core.Repository[entities.Expense]
}
