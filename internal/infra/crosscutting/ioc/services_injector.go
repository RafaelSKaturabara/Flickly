package inversion_of_control

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/cross-cutting/utilities"
	infraRepositories "flickly/internal/infra/data/users/repositories"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()

	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infraRepositories.NewUserRepository())
}
