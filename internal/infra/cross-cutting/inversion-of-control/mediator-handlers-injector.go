package inversion_of_control

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/commands"
	"flickly/internal/infra/cross-cutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)

	mediatR.Register("CreateUserCommand", commands.NewCreateUserCommandHandler(serviceCollection))
}
