package ioc

import (
	"github.com/rkaturabara/flickly/internal/application/commons/auto_mapper"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

func InitAutomapper(serviceCollection utilities.IServiceCollection) {
	automapper := utilities.NewAutoMapper()
	utilities.AddService[utilities.Mapper](serviceCollection, automapper)
	auto_mapper.ViewModelAutomapperConfig(serviceCollection)
}
