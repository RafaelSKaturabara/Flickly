package auto_mapper

import (
	"reflect"

	"github.com/rkaturabara/flickly/internal/api/commons/view_model"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
)

func ViewModelAutomapperConfig(serviceCollection utilities.IServiceCollection) {
	mapper := utilities.GetService[utilities.Mapper](serviceCollection)

	mapper.AddMapping(
		reflect.TypeOf(core.DomainError{}),
		reflect.TypeOf(view_model.ErrorResponse{}),
		func(source, dest reflect.Value) error {
			dest.FieldByName("Code").Set(source.FieldByName("Code"))
			dest.FieldByName("Message").Set(source.FieldByName("Message"))

			errorMethod := source.MethodByName("Error")
			if errorMethod.IsValid() {
				errorResult := errorMethod.Call(nil)
				if len(errorResult) > 0 {
					dest.FieldByName("InternalMessage").Set(errorResult[0])
				}
			}

			return nil
		},
	)
}
