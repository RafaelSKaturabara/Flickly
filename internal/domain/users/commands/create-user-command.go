package commands

import (
	"flickly/internal/domain/core"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	inversionofcontrol "flickly/internal/infra/cross-cutting/utilities"
	"github.com/gin-gonic/gin"
)

type CreateUserCommand struct {
	mediator.Request
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserCommandHandler struct {
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
}

func NewCreateUserCommandHandler(serviceCollection inversionofcontrol.IServiceCollection) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		mediator:       inversionofcontrol.GetService[mediator.Mediator](serviceCollection),
		userRepository: inversionofcontrol.GetService[repositories.IUserRepository](serviceCollection),
	}
}

func (h *CreateUserCommandHandler) Handle(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateUserCommand)
	user := entities.NewUser(command.Name, command.Email)
	err := h.userRepository.CreateUser(user)
	if err != nil {
		return nil, core.ErrUserAlreadyExist(err)
	}
	return user, nil
}
