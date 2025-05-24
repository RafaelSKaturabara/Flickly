package ioc

import (
	"reflect"

	"github.com/rkaturabara/flickly/internal/api/commons/auto_mapper"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

func InitAutomapper(serviceCollection utilities.IServiceCollection) {
	automapper := utilities.NewAutoMapper()
	serviceCollection.AddServiceInstance(reflect.TypeOf((*utilities.Mapper)(nil)).Elem(), automapper)
	auto_mapper.ViewModelAutomapperConfig(serviceCollection)
}
