package ioc

import (
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/command_handlers"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)
	mediatR.Register("CreateUserCommand", command_handlers.NewCreateUserCommandHandler(serviceCollection))
	mediatR.Register("CreateTokenCommand", command_handlers.NewCreateTokenCommandHandler(serviceCollection))
}
