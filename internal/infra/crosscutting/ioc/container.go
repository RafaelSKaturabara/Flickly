package ioc

import (
	"reflect"

	domainrepos "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	userrepos "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/repositories"
	users "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity"
)

type Container struct {
	services utilities.IServiceCollection
}

func NewContainer() *Container {
	container := &Container{
		services: utilities.NewServiceCollection(),
	}

	// Registrar o repositório de usuários
	userRepo := userrepos.NewUserRepository(users.GetUsersLocalDBConnection())
	container.services.AddServiceInstance(reflect.TypeOf((*domainrepos.IUserRepository)(nil)).Elem(), userRepo)

	return container
}

func (c *Container) GetUserRepository() domainrepos.IUserRepository {
	return c.services.GetServiceByType(reflect.TypeOf((*domainrepos.IUserRepository)(nil)).Elem()).(domainrepos.IUserRepository)
}
