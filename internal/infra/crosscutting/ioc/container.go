package ioc

import (
	"reflect"

	domainrepos "github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
	userrepos "github.com/rkaturabara/flickly/internal/infra/data/users/repositories"
)

type Container struct {
	services utilities.IServiceCollection
}

func NewContainer() *Container {
	container := &Container{
		services: utilities.NewServiceCollection(),
	}

	// Registrar o repositório de usuários
	userRepo := userrepos.NewUserRepository()
	container.services.AddServiceInstance(reflect.TypeOf((*domainrepos.IUserRepository)(nil)).Elem(), userRepo)

	return container
}

func (c *Container) GetUserRepository() domainrepos.IUserRepository {
	return c.services.GetServiceByType(reflect.TypeOf((*domainrepos.IUserRepository)(nil)).Elem()).(domainrepos.IUserRepository)
}
