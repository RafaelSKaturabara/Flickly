package ioc

import (
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/commands"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)
	mediatR.Register("CreateUserCommand", commands.NewCreateUserCommandHandler(serviceCollection))
}
