package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IncomeDB struct {
	entities.Income `gorm:"embedded"`
	UserSimpleRuleID          uuid.UUID        `gorm:"column:user_id;not null"`
	User            UserSimpleRuleDB `gorm:"foreignKey:UserID;references:ID"`
}

func (IncomeDB) TableName() string {
	return "incomes"
}

func (db *IncomeDB) ToDomain() entities.Income {
	return db.Income
}

func (db *IncomeDB) FromDomain(d entities.Income) {
	db.Income = d
}

func SeedIncome(db *gorm.DB) {}
