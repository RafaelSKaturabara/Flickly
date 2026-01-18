package ioc

import (
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
	infrarepositories "github.com/rkaturabara/flickly/internal/infra/data/users/repositories"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infrarepositories.NewUserRepository())
}
