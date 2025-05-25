package command_handlers

import (
	"context"

	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/domain/users/services"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

type RefreshTokenCommand struct {
	RefreshToken string
}

type RefreshTokenCommandHandler struct {
	mediator       mediator.Mediator
	userRepository repositories.IUserRepository
}

func NewRefreshTokenCommandHandler(serviceCollection utilities.IServiceCollection) *RefreshTokenCommandHandler {
	return &RefreshTokenCommandHandler{
		mediator:       utilities.GetService[mediator.Mediator](serviceCollection),
		userRepository: utilities.GetService[repositories.IUserRepository](serviceCollection),
	}
}

func (h *RefreshTokenCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(RefreshTokenCommand)

	// 1. Validar o refresh token e obter o user_id
	validateService := services.NewValidateRefreshTokenService()
	userID, err := validateService.ValidateRefreshToken(command.RefreshToken)
	if err != nil {
		return nil, core.ErrInvalidToken(err)
	}

	// 2. Buscar o usuário
	user, err := h.userRepository.GetUserByID(c, userID)
	if err != nil || user == nil {
		return nil, core.ErrInvalidToken(err)
	}

	// 3. Verificar se o refresh token ainda é válido
	if user.RefreshToken != command.RefreshToken {
		return nil, core.ErrInvalidToken(nil)
	}

	// 4. Gerar novo access token
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
