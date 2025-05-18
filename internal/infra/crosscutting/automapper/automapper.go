package automapper

import (
	"reflect"
)

// Mapper function to map source to destination
func Map(source, dest interface{}) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(dest).Elem()

	// Mapear campos comuns
	fields := []string{"ID", "Email", "CreatedAt", "Age", "Address"}
	for _, field := range fields {
		sourceField := sourceValue.FieldByName(field)
		destField := destValue.FieldByName(field)
		if sourceField.IsValid() && destField.IsValid() {
			destField.Set(sourceField)
		}
	}

	// Mapear Name para FullName
	sourceName := sourceValue.FieldByName("Name")
	destFullName := destValue.FieldByName("FullName")
	if sourceName.IsValid() && destFullName.IsValid() {
		destFullName.Set(sourceName)
	}

	// Mapear mensagem de erro
	errorMethod := sourceValue.MethodByName("Error")
	if errorMethod.IsValid() && errorMethod.CanInterface() {
		errorResult := errorMethod.Call(nil)
		if len(errorResult) > 0 {
			destMessage := destValue.FieldByName("InternalMessage")
			if destMessage.IsValid() {
				destMessage.Set(errorResult[0])
			}
		}
	}
	return nil
}
