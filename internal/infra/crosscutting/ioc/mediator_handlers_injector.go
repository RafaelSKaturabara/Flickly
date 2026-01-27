package ioc

import (
	simpleRuleCommands "github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/commands"
	simpleRuleQueries "github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/queries"
	command_handlers "github.com/RafaelSKaturabara/Flickly/internal/application/users/handlers"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)
	mediatR.Register("CreateUserCommand", command_handlers.NewCreateUserCommandHandler(serviceCollection))
	mediatR.Register("CreateTokenCommand", command_handlers.NewCreateTokenCommandHandler(serviceCollection))

	// Simple Rule Commands
	mediatR.Register("CreateSimpleRuleCommand", simpleRuleCommands.NewCreateExpenseCommandHandler(serviceCollection))

	// Simple Rule Queries
	mediatR.Register("ListAllSubcategoriesQuery", simpleRuleQueries.NewListAllSubcategoriesQueryHandler(serviceCollection))
}
