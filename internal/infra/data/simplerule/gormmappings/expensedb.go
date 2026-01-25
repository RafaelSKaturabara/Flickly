package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ExpenseDB struct {
	entities.Expense `gorm:"embedded"`
	UserID           uuid.UUID        `gorm:"column:user_id;not null"`
	User             SimpleRuleUserDB `gorm:"foreignKey:UserID;references:ID"`
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
