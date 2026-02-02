package commands

import (
	"context"
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/google/uuid"
)

type GetAuthorizeCommand struct {
	Challenge     string
	ClientID uuid.UUID
}

type GetAuthorizeCommandHandler struct {
	mediator       mediator.Mediator
	clientRepository repositories.IClientRepository
	authCodeRepository repositories.IAuthCodeRepository
}

func NewGetAuthorizeCommandHandler(serviceCollection utilities.IServiceCollection) *GetAuthorizeCommandHandler {
	return &GetAuthorizeCommandHandler{
		mediator:       utilities.GetService[mediator.Mediator](serviceCollection),
		clientRepository: utilities.GetService[repositories.IClientRepository](serviceCollection),
		authCodeRepository: utilities.GetService[repositories.IAuthCodeRepository](serviceCollection),
	}
}

func (h *GetAuthorizeCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(GetAuthorizeCommand)

	// Gerar code
	code := uuid.New().String()

	// Get Default UserAccess
	client, err := h.clientRepository.GetByID(command.ClientID)
	if err != nil {
		return nil, err
	}

	authCode := entities.NewAuthCode(code, client.ID, nil, command.Challenge, time.Now().Add(time.Minute * 600))

	err = h.authCodeRepository.Create(authCode)
	if err != nil {
		return nil, err
	}

	authCode.Client = client

	return authCode, nil
}
