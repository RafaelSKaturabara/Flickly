package commandhandlers

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
	"github.com/rkaturabara/flickly/internal/domain/simplerule/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
	"github.com/shopspring/decimal"
)

type CreateExpenseCommandHandler struct {
	mediator          mediator.Mediator
	expenseRepository repositories.IExpenseRepository
}

func NewCreateExpenseCommandHandler(serviceCollection utilities.IServiceCollection) *CreateExpenseCommandHandler {
	return &CreateExpenseCommandHandler{
		mediator:          utilities.GetService[mediator.Mediator](serviceCollection),
		expenseRepository: utilities.GetService[repositories.IExpenseRepository](serviceCollection),
	}
}

func (h *CreateExpenseCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateRouteCommand)

	expense := entities.NewExpense(command.UserID, command.CommitmentID, command.Date, command.SubcategoryID, command.Amount, command.Description, command.InstallmentNumber, command.IsReference)

	if !expense.IsValid() {
		return nil, core.InvalidEntityError(expense.GetErrors())
	}

	err := h.expenseRepository.Create(*expense)
	if err != nil {
		return nil, core.DoNotUseThisGenericError(err)
	}

	return expense, nil
}

type CreateRouteCommand struct {
	UserID            uuid.UUID
	CommitmentID      *uuid.UUID
	Date              time.Time
	SubcategoryID     uuid.UUID
	Amount            decimal.Decimal
	Description       string
	InstallmentNumber *int
	IsReference       bool
}
