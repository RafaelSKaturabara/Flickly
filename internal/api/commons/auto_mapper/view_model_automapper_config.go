package auto_mapper

import (
	"flickly/internal/api/commons/view_model"
	"flickly/internal/domain/core"
	"flickly/internal/infra/crosscutting/utilities"
	"reflect"
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
