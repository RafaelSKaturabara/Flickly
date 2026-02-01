package auto_mapper

import (
	"reflect"
	"strings"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/commons/view_model"
	"github.com/RafaelSKaturabara/Flickly/internal/presentation/identity/viewmodel"
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

	mapper.AddMapping(
		reflect.TypeOf(entities.User{}),
		reflect.TypeOf(viewmodel.TokenResponse{}),
		func(source, dest reflect.Value) error {
			dest.FieldByName("AccessToken").Set(source.FieldByName("AccessToken"))
			dest.FieldByName("TokenType").Set(reflect.ValueOf("Bearer"))
			dest.FieldByName("ExpiresIn").Set(source.FieldByName("TokenExpiry"))

			// Converte o array de TokenScopes em uma string separada por vírgula
			tokenScopes := source.FieldByName("TokenScopes").Interface().([]string)
			scopesString := strings.Join(tokenScopes, ",")
			dest.FieldByName("Scope").Set(reflect.ValueOf(scopesString))

			return nil
		},
	)
}
