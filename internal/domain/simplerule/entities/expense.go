package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/shopspring/decimal"
)

type Expense struct {
	core.BaseEntity
	UserID            uuid.UUID
	CommitmentID      *uuid.UUID      // Refere-se ao Commitment
	Date              time.Time       // Data do lançamento
	SubcategoryID     uuid.UUID       // Ex: Shelter, Entertainment (o "SubType")
	Amount            decimal.Decimal // Valor em centavos para evitar erros de float
	Description       string          // Descrição adicional
	InstallmentNumber *int            // Controle de parcelas (ex: 6). O 6/12 está no commitment
	IsReference       bool            // Campo "É apenas referência?". Utilizado para simular e contabilizar despesas futuras
}

// NewExpense é o construtor (Factory) para a entidade Expense
func NewExpense(
	userID uuid.UUID,
	commitmentID *uuid.UUID,
	date time.Time,
	subcategoryID uuid.UUID,
	amount decimal.Decimal,
	description string,
	installmentNumber *int,
	isReference bool,
) *Expense {
	return &Expense{
		BaseEntity:        core.NewBaseEntity(),
		Description:       description,
		Amount:            amount,
		SubcategoryID:     subcategoryID,
		Date:              date,
		CommitmentID:      commitmentID,
		InstallmentNumber: installmentNumber,
		IsReference:       isReference,
	}
}

func (u *Expense) IsValid() bool {
	return true
}
