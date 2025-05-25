package utilities

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockInterface é uma interface de exemplo para testes
type MockInterface interface {
	DoSomething() string
}

// MockImplementation é uma implementação de exemplo para testes
type MockImplementation struct{}

func (m *MockImplementation) DoSomething() string {
	return "test"
}

// MockIncorrectImplementation não implementa completamente a interface MockInterface
type MockIncorrectImplementation struct{}

type TestService struct {
	Value string
}

type TestInterface interface {
	GetValue() string
}

func (s *TestService) GetValue() string {
	return s.Value
}

func TestNewServiceCollection(t *testing.T) {
	// Execução
	collection := NewServiceCollection()

	// Verificações
	assert.NotNil(t, collection, "NewServiceCollection deve retornar uma instância não nula")
}

func TestAddServiceInstance(t *testing.T) {
	// Configuração
	collection := NewServiceCollection()
	service := &TestService{Value: "test"}
	typeOfInterface := reflect.TypeOf((*TestInterface)(nil)).Elem()

	// Execução
	result := collection.AddServiceInstance(typeOfInterface, service)

	// Verificações
	assert.NotNil(t, result, "AddServiceInstance deve retornar uma instância não nula")
}

func TestGetServiceByType(t *testing.T) {
	// Configuração
	collection := NewServiceCollection()
	service := &TestService{Value: "test"}
	typeOfInterface := reflect.TypeOf((*TestInterface)(nil)).Elem()
	collection.AddServiceInstance(typeOfInterface, service)

	// Execução - serviço existente
	retrievedService := collection.GetServiceByType(typeOfInterface)

	// Verificações
	assert.NotNil(t, retrievedService, "Deve retornar o serviço quando encontrado")
	assert.Equal(t, service, retrievedService, "O serviço retornado deve ser o mesmo que foi adicionado")

	// Execução - serviço não existente
	type OtherInterface interface{}
	retrievedService = collection.GetServiceByType(reflect.TypeOf((*OtherInterface)(nil)).Elem())

	// Verificações
	assert.Nil(t, retrievedService, "Deve retornar nil para serviço não encontrado")
}

func TestAddService(t *testing.T) {
	// Configuração
	collection := NewServiceCollection()

	// Execução
	result := AddService[TestInterface](collection, &TestService{Value: "test"})

	// Verificações
	assert.NotNil(t, result, "AddService deve retornar uma instância não nula")
}

func TestGetService(t *testing.T) {
	// Configuração
	collection := NewServiceCollection()
	service := &TestService{Value: "test"}
	AddService[TestInterface](collection, service)

	// Execução - serviço existente
	retrievedService := GetService[TestInterface](collection)

	// Verificações
	assert.NotNil(t, retrievedService, "Deve retornar o serviço quando encontrado")
	assert.Equal(t, service, retrievedService, "O serviço retornado deve ser o mesmo que foi adicionado")

	// Execução - serviço não existente
	type OtherInterface interface{}
	retrievedOtherService := GetService[OtherInterface](collection)

	// Verificações
	assert.Nil(t, retrievedOtherService, "Deve retornar nil para serviço não encontrado")
}

func TestServiceCollection_ConcurrentAccess(t *testing.T) {
	// Configuração
	collection := NewServiceCollection()
	service := &TestService{Value: "test"}

	// Execução - adiciona o serviço
	result := AddService[TestInterface](collection, service)
	assert.NotNil(t, result)

	// Execução - tenta adicionar o mesmo serviço novamente
	result = AddService[TestInterface](collection, service)
	assert.NotNil(t, result)

	// Execução - tenta adicionar um serviço diferente da mesma interface
	result = AddService[TestInterface](collection, &TestService{Value: "another"})
	assert.NotNil(t, result)
}
