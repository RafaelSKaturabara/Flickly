package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseDB struct {
	core.BaseEntityDB `gorm:"embedded"`
	entities.Expense  `gorm:"embedded"`
	UserSimpleRuleID            uuid.UUID        `gorm:"column:user_id;not null"`
	User              UserSimpleRuleDB `gorm:"foreignKey:UserID;references:ID"`
}

func (ExpenseDB) TableName() string {
	return "expenses"
}

func (db *ExpenseDB) ToDomain() entities.Expense {
	return db.Expense
}

func (db *ExpenseDB) FromDomain(d entities.Expense) {
	db.Expense = d
}

func SeedExpense(db *gorm.DB) {}
