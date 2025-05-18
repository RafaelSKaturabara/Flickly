package commands

import (
	"flickly/internal/domain/core"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
)

type CreateUserCommand struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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

func (h *CreateUserCommandHandler) Handle(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateUserCommand)
	user := entities.NewUser(command.Name, command.Email)
	err := h.userRepository.CreateUser(user)
	if err != nil {
		return nil, core.ErrUserAlreadyExist(err)
	}
	return user, nil
}
