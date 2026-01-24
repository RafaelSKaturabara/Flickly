package ioc

import (
	simpleRuleCommandHandlers "github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/handlers"
	command_handlers "github.com/RafaelSKaturabara/Flickly/internal/application/users/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)
	mediatR.Register("CreateUserCommand", command_handlers.NewCreateUserCommandHandler(serviceCollection))
	mediatR.Register("CreateTokenCommand", command_handlers.NewCreateTokenCommandHandler(serviceCollection))
	mediatR.Register("CreateSimpleRuleCommand", simpleRuleCommandHandlers.NewCreateExpenseCommandHandler(serviceCollection))
}
