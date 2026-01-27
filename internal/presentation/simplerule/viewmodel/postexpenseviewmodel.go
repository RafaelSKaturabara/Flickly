package viewmodel

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PostExpenseRequest struct {
	CommitmentID      *uuid.UUID      `json:"commitmentId"`      // Refere-se ao Commitment
	Date              time.Time       `json:"date"`              // Data do lançamento
	SubcategoryID     uuid.UUID       `json:"subcategoryId"`     // Ex: Shelter, Entertainment (o "SubType")
	Amount            decimal.Decimal `json:"amount"`            // Valor em centavos para evitar erros de float
	Description       string          `json:"description"`       // Descrição adicional
	InstallmentNumber *int            `json:"installmentNumber"` // Controle de parcelas (ex: 6). O 6/12 está no commitment
	IsReference       bool            `json:"isReference"`       // Campo "É apenas referência?". Utilizado para simular e contabilizar despesas futuras
}


type PostExpenseResponse struct {
	ID	uuid.UUID
}
