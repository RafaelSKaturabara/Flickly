package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
)

type IIncomeRepository interface {
	core.Repository[entities.Income]
}
