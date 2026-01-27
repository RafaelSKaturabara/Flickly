package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type UserSimpleRuleRepository struct {
	core.GormRepository[entities.UserSimpleRule, gormmappings.UserSimpleRuleDB]
}

func NewUserSimpleRuleRepository(db *gorm.DB) *UserSimpleRuleRepository {
	return &UserSimpleRuleRepository{
		core.GormRepository[entities.UserSimpleRule, gormmappings.UserSimpleRuleDB](*core.NewGormRepository[entities.UserSimpleRule, gormmappings.UserSimpleRuleDB](db)),
	}
}
