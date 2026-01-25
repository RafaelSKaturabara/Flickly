package ioc

import (
	"reflect"

	domainrepos "github.com/RafaelSKaturabara/Flickly/internal/domain/users/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	dbinit "github.com/RafaelSKaturabara/Flickly/internal/infra/data/init"
	userrepos "github.com/RafaelSKaturabara/Flickly/internal/infra/data/users/repositories"
)

type Container struct {
	services utilities.IServiceCollection
}

func NewContainer() *Container {
	container := &Container{
		services: utilities.NewServiceCollection(),
	}

	// Registrar o repositório de usuários
	userRepo := userrepos.NewUserRepository(dbinit.GetUsersLocalDBConnection())
	container.services.AddServiceInstance(reflect.TypeOf((*domainrepos.IUserRepository)(nil)).Elem(), userRepo)

	return container
}

func (c *Container) GetUserRepository() domainrepos.IUserRepository {
	return c.services.GetServiceByType(reflect.TypeOf((*domainrepos.IUserRepository)(nil)).Elem()).(domainrepos.IUserRepository)
}
