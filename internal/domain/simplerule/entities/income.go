package entities

import (
	"github.com/google/uuid"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/shopspring/decimal"
)

type Income struct {
	core.BaseEntity
	UserID     uuid.UUID
	Amount     decimal.Decimal
	Source     string
	IsAfterTax bool
}

func (u *Income) IsValid() bool {
	return true
}
