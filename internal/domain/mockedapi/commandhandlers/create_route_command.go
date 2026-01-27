package commandhandlers

import (
	"context"

	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/mockedapi/entities"
	"github.com/rkaturabara/flickly/internal/domain/mockedapi/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

type CreateRouteCommandHandler struct {
	mediator       mediator.Mediator
	routeRepository repositories.IRouteRepository
}

func NewCreateRouteCommandHandler(serviceCollection utilities.IServiceCollection) *CreateRouteCommandHandler {
	return &CreateRouteCommandHandler{
		mediator:       utilities.GetService[mediator.Mediator](serviceCollection),
		routeRepository: utilities.GetService[repositories.IRouteRepository](serviceCollection),
	}
}

func (h *CreateRouteCommandHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	command := request.(CreateRouteCommand)

	route := entities.NewRoute(command.Paths, command.Methods, command.Mocks)
	err := h.routeRepository.CreateRoute(c, route)
	if err != nil {
		return nil, core.DoNotUseThisGenericError(err)
	}
	return route, nil
}
	

type CreateRouteCommand struct {
	Paths   string
	Methods entities.HttpMethod
	Mocks   []*entities.Mock
}