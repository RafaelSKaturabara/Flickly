package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type UserSimpleRuleDB struct {
	core.BaseEntityDB       `gorm:"embedded"`
	entities.UserSimpleRule `gorm:"embedded"`
	Nickname                string `gorm:"type:varchar(100);not null"`
}

func (UserSimpleRuleDB) TableName() string {
	return "usersimplerule"
}

func (db *UserSimpleRuleDB) ToDomain() entities.UserSimpleRule {
	return db.UserSimpleRule
}

func (db *UserSimpleRuleDB) FromDomain(d entities.UserSimpleRule) {
	db.UserSimpleRule = d
}

func SeedUserSimpleRule(db *gorm.DB) {}