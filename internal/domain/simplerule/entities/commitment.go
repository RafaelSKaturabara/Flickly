package entities

import (
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Commitment struct {
	core.BaseEntity
	Description       string          
	TotalAmount       decimal.Decimal
	TotalInstallments int             
	StartDate         time.Time       
	IsActive          bool            
	UserID            uuid.UUID
	User              SimpleRuleUser
}

func (c *Commitment) GetRemainingInstallments(currentInstallment int) int {
	if currentInstallment >= c.TotalInstallments {
		return 0
	}
	return c.TotalInstallments - currentInstallment
}

func (u *Commitment) IsValid() bool {
	return true
}
