package command_handlers

import (
	"context"

	"github.com/rkaturabara/flickly/internal/application/commons/middleware"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/domain/users/services"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

type CreateTokenCommand struct {
	GrantType    string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	Scope        string
}

type CreateCreateTokenCommandHandler struct {
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
}

func NewCreateTokenCommandHandler(serviceCollection utilities.IServiceCollection) *CreateCreateTokenCommandHandler {
	return &CreateCreateTokenCommandHandler{
		mediator:       utilities.GetService[mediator.Mediator](serviceCollection),
		userRepository: utilities.GetService[repositories.IUserRepository](serviceCollection),
	}
}

func (h *CreateCreateTokenCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateTokenCommand)

	var err error
	var user *entities.User

	if command.GrantType == "password" {
		// Busca o usuário pelo email
		user, err = h.userRepository.GetUserByEmailAndPasswordAndClientAndSecret(c, command.Username, command.Password, command.ClientID, command.ClientSecret)
		if err != nil || user == nil {
			// corrigir erro na consulta
			return nil, core.ErrInvalidCredentials(err)
		}
	} else if command.GrantType == "refresh_token" {
		userJwt, ok := c.Value(middleware.UserContextKey).(*entities.User)
		user, err = h.userRepository.GetUserByID(c, userJwt.GetID())
		if !ok || user == nil {
			// corrigir erro na consulta
			return nil, core.ErrInvalidCredentials(err)
		}
	}

	var serviceRules []core.Service
	serviceRules = append(serviceRules, services.NewGenerateJWTService())

	for _, service := range serviceRules {
		if service.AbleToRun(c, user) {
			err = service.Run(c, user)
			if err != nil {
				return nil, err
			}
		}
	}

	return user, nil
}
