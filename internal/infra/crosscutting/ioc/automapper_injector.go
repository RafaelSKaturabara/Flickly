package ioc

import (
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/auto_mapper"
)

func InitAutomapper(serviceCollection utilities.IServiceCollection) {
	automapper := utilities.NewAutoMapper()
	utilities.AddService[utilities.Mapper](serviceCollection, automapper)
	auto_mapper.ViewModelAutomapperConfig(serviceCollection)
}
