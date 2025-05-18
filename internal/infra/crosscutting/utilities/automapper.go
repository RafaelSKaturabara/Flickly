package utilities

import (
	"errors"
	"reflect"
)

// Mapper é a interface principal do automapper
type Mapper interface {
	Map(source, destination interface{}) error
	AddMapping(sourceType, destType reflect.Type, mapping func(source, destination reflect.Value) error)
	MapSlice(source, destination interface{}) error
}

// AutoMapper é a implementação padrão do Mapper
type AutoMapper struct {
	mappings map[string]func(source, destination reflect.Value) error
}

// NewAutoMapper cria uma nova instância do AutoMapper
func NewAutoMapper() Mapper {
	return &AutoMapper{
		mappings: make(map[string]func(source, destination reflect.Value) error),
	}
}

// Map mapeia os campos de source para destination
func (m *AutoMapper) Map(source, destination interface{}) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(destination)

	// Verificar se destination é um ponteiro
	if destValue.Kind() != reflect.Ptr {
		return errors.New("destination must be a pointer")
	}

	// Obter o valor apontado
	destValue = destValue.Elem()

	// Verificar se source é um ponteiro
	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}

	// Obter os tipos
	sourceType := sourceValue.Type()
	destType := destValue.Type()

	// Criar chave para o mapeamento
	key := sourceType.String() + "->" + destType.String()

	// Verificar se já existe um mapeamento personalizado
	if mapping, exists := m.mappings[key]; exists {
		return mapping(sourceValue, destValue)
	}

	// Mapeamento padrão
	return m.defaultMapping(sourceValue, destValue)
}

// defaultMapping implementa o mapeamento padrão entre structs
func (m *AutoMapper) defaultMapping(source, destination reflect.Value) error {
	// Verificar se ambos são structs
	if source.Kind() != reflect.Struct || destination.Kind() != reflect.Struct {
		return errors.New("both source and destination must be structs")
	}

	// Iterar sobre os campos da struct de destino
	for i := 0; i < destination.NumField(); i++ {
		destField := destination.Type().Field(i)
		sourceField := source.FieldByName(destField.Name)

		// Se o campo existe na source
		if sourceField.IsValid() {
			// Se os tipos são compatíveis
			if sourceField.Type().AssignableTo(destField.Type) {
				destination.Field(i).Set(sourceField)
			}
		}
	}

	return nil
}

// AddMapping adiciona um mapeamento personalizado
func (m *AutoMapper) AddMapping(sourceType, destType reflect.Type, mapping func(source, destination reflect.Value) error) {
	key := sourceType.String() + "->" + destType.String()
	m.mappings[key] = mapping
}

// MapSlice mapeia um slice de source para um slice de destination
func (m *AutoMapper) MapSlice(source, destination interface{}) error {
	sourceValue := reflect.ValueOf(source)
	destValue := reflect.ValueOf(destination)

	// Verificar se destination é um ponteiro para slice
	if destValue.Kind() != reflect.Ptr || destValue.Elem().Kind() != reflect.Slice {
		return errors.New("destination must be a pointer to slice")
	}

	// Verificar se source é um slice
	if sourceValue.Kind() != reflect.Slice {
		return errors.New("source must be a slice")
	}

	// Obter o slice de destino
	destSlice := destValue.Elem()

	// Criar um novo slice com a capacidade correta
	newSlice := reflect.MakeSlice(destSlice.Type(), sourceValue.Len(), sourceValue.Cap())

	// Mapear cada elemento
	for i := 0; i < sourceValue.Len(); i++ {
		// Criar uma nova instância do tipo de destino
		newItem := reflect.New(destSlice.Type().Elem()).Elem()

		// Mapear o elemento
		if err := m.Map(sourceValue.Index(i).Interface(), newItem.Addr().Interface()); err != nil {
			return err
		}

		// Adicionar ao novo slice
		newSlice.Index(i).Set(newItem)
	}

	// Atualizar o slice de destino
	destSlice.Set(newSlice)

	return nil
}
