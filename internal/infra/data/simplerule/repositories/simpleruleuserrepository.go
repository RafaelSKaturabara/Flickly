package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type SimpleRuleUserRepository struct {
	core.GormRepository[entities.SimpleRuleUser, gormmappings.SimpleRuleUserDB]
}

func NewSimpleRuleUserRepository(db *gorm.DB) *SimpleRuleUserRepository {
	return &SimpleRuleUserRepository{
		core.GormRepository[entities.SimpleRuleUser, gormmappings.SimpleRuleUserDB](*core.NewGormRepository[entities.SimpleRuleUser, gormmappings.SimpleRuleUserDB](db)),
	}
}
