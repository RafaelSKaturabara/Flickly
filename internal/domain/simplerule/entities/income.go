package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Income struct {
	core.BaseEntity
	Amount           decimal.Decimal
	Source           string
	IsAfterTax       bool
	UserSimpleRuleID uuid.UUID
	User             UserSimpleRule
}

func NewIncome(amount decimal.Decimal, source string, isAfterTax bool, userSimpleRuleID uuid.UUID) Income {
	return Income{
		BaseEntity:       core.NewBaseEntity(),
		Amount:           amount,
		Source:           source,
		IsAfterTax:       isAfterTax,
		UserSimpleRuleID: userSimpleRuleID,
	}
}

func (u *Income) IsValid() bool {
	return true
}
