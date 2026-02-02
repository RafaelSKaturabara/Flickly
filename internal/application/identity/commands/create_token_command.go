package commands

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/application"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/services"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/valueobjects"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
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
		passwordHash, err := valueobjects.NewPasswordHash(command.Password)
		if err != nil {
			return nil, err
		}

		user, err = h.userRepository.GetUserByEmailAndPasswordAndClient(c, command.Username, passwordHash.GetHash(), command.ClientID)
		if err != nil || user == nil {
			// corrigir erro na consulta
			return nil, core.ErrInvalidCredentials(err)
		}
	} else if command.GrantType == "refresh_token" {
		userJwt, ok := c.Value(application.UserContextKey).(*entities.User)
		user, err = h.userRepository.GetByID(userJwt.GetID())
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
