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
	UserSimpleRuleID  uuid.UUID
	User              UserSimpleRule
}

func NewCommitment(description string, totalAmount decimal.Decimal, totalInstallments int, startDate time.Time, userSimpleRuleID uuid.UUID) Commitment {
	return Commitment{
		BaseEntity:        core.NewBaseEntity(),
		Description:       description,
		TotalAmount:       totalAmount,
		TotalInstallments: totalInstallments,
		StartDate:         startDate,
		IsActive:          true,
		UserSimpleRuleID:  userSimpleRuleID,
	}
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
