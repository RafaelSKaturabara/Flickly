package entities

import (
	"time"

	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/shopspring/decimal"
)

type Commitment struct {
	core.BaseEntity
	Description       string          `json:"description"`
	TotalAmount       decimal.Decimal `json:"total_amount"`
	TotalInstallments int             `json:"total_installments"`
	StartDate         time.Time       `json:"start_date"`
	IsActive          bool            `json:"is_active"`
}

// GetRemainingInstallments é um método auxiliar para calcular quantas parcelas faltam,
// dado o número da parcela atual paga.
func (c *Commitment) GetRemainingInstallments(currentInstallment int) int {
	if currentInstallment >= c.TotalInstallments {
		return 0
	}
	return c.TotalInstallments - currentInstallment
}

func (u *Commitment) IsValid() bool {
	return true
}
