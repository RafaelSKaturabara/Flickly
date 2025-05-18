package ioc

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/commands"
	"flickly/internal/infra/crosscutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)

	mediatR.Register("CreateUserCommand", commands.NewCreateUserCommandHandler(serviceCollection))
}
