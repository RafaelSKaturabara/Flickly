package ioc

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	infrarepositories "flickly/internal/infra/data/users/repositories"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infrarepositories.NewUserRepository())
}
