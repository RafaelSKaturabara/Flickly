package entities

import (
	"github.com/google/uuid"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/shopspring/decimal"
)

type Income struct {
	core.BaseEntity
	Amount     decimal.Decimal
	Source     string
	IsAfterTax bool
	UserID     uuid.UUID
	User       SimpleRuleUser
}

func (u *Income) IsValid() bool {
	return true
}
