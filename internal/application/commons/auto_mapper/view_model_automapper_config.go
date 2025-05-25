package auto_mapper

import (
	"reflect"
	"strings"

	"github.com/rkaturabara/flickly/internal/application/commons/view_model"
	"github.com/rkaturabara/flickly/internal/application/users/viewmodel"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
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
