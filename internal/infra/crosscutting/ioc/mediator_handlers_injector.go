package ioc

import (
	simpleRuleCommands "github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/commands"
	simpleRuleQueries "github.com/RafaelSKaturabara/Flickly/internal/application/simplerule/queries"
	identityCommands "github.com/RafaelSKaturabara/Flickly/internal/application/identity/commands"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
)

func InjectMediatorHandlers(serviceCollection utilities.IServiceCollection) {
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)

	// Identity Commands
	mediatR.Register("CreateUserCommand", identityCommands.NewCreateUserCommandHandler(serviceCollection))
	mediatR.Register("CreateTokenCommand", identityCommands.NewCreateTokenCommandHandler(serviceCollection))
	mediatR.Register("CreateLoginCommand", identityCommands.NewCreateLoginCommandHandler(serviceCollection))
	mediatR.Register("GetAuthorizeCommand", identityCommands.NewGetAuthorizeCommandHandler(serviceCollection))

	// Simple Rule Commands
	mediatR.Register("CreateSimpleRuleCommand", simpleRuleCommands.NewCreateExpenseCommandHandler(serviceCollection))

	// Simple Rule Queries
	mediatR.Register("ListAllSubcategoriesQuery", simpleRuleQueries.NewListAllSubcategoriesQueryHandler(serviceCollection))
}
