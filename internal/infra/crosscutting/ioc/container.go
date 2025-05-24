package ioc

import (
	"reflect"

	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

type Container struct {
	services utilities.IServiceCollection
}

func NewContainer() *Container {
	return &Container{
		services: utilities.NewServiceCollection(),
	}
}

func (c *Container) GetUserRepository() repositories.IUserRepository {
	return c.services.GetServiceByType(reflect.TypeOf((*repositories.IUserRepository)(nil)).Elem()).(repositories.IUserRepository)
}
