package command_handlers

import (
	"context"

	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

type CreateUserCommand struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Senha não é serializada em JSON
}

type CreateUserCommandHandler struct {
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
}

func NewCreateUserCommandHandler(serviceCollection utilities.IServiceCollection) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		mediator:       utilities.GetService[mediator.Mediator](serviceCollection),
		userRepository: utilities.GetService[repositories.IUserRepository](serviceCollection),
	}
}

func (h *CreateUserCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateUserCommand)
	password, err := utilities.Encrypt(command.Password)
	if err != nil {
		return nil, err
	}

	user := entities.NewUser(command.Name, command.Email, "ClientID", "ClientSecret", password)
	err = h.userRepository.CreateUser(c, user)
	if err != nil {
		return nil, core.ErrUserAlreadyExist(err)
	}
	return user, nil
}
