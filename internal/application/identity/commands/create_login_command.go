package commands

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/valueobjects"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/google/uuid"
)

type CreateLoginCommand struct {
	Name     string
	Email    string
	Password string // Senha não é serializada em JSON
	ClientID uuid.UUID
}

type CreateLoginCommandHandler struct {
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
}

func NewCreateLoginCommandHandler(serviceCollection utilities.IServiceCollection) *CreateLoginCommandHandler {
	return &CreateLoginCommandHandler{
		mediator:       utilities.GetService[mediator.Mediator](serviceCollection),
		userRepository: utilities.GetService[repositories.IUserRepository](serviceCollection),
	}
}

func (h *CreateLoginCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateLoginCommand)
	passwordHash, err := valueobjects.NewPasswordHash(command.Password)
	if err != nil {
		return nil, err	
	}

	// Get Default UserAccess
	panic("not implemented UserAccess")

	user := entities.NewUser(command.Name, command.Email, passwordHash, nil)
	err = h.userRepository.Create(user)
	if err != nil {
		return nil, core.ErrUserAlreadyExist(err)
	}
	return user, nil
}
