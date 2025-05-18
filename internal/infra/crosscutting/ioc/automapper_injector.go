package ioc

import (
	"flickly/internal/api/commons/auto_mapper"
	"flickly/internal/infra/crosscutting/utilities"
)

func InitAutomapper(serviceCollection utilities.IServiceCollection) {
	automapper := utilities.NewAutoMapper()

	utilities.AddService[utilities.Mapper](serviceCollection, automapper)

	auto_mapper.ViewModelAutomapperConfig(serviceCollection)
}
