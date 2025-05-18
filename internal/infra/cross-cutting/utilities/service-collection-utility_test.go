package utilities

import (
	"reflect"
	"testing"
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

func TestNewServiceCollection(t *testing.T) {
	serviceCollection := NewServiceCollection()
	if serviceCollection == nil {
		t.Fatal("NewServiceCollection deve retornar uma instância não nula")
	}
}

func TestAddServiceInstance(t *testing.T) {
	serviceCollection := NewServiceCollection()
	implementation := &MockImplementation{}
	typeOfInterface := reflect.TypeOf((*MockInterface)(nil)).Elem()

	// Testando adição válida
	result := serviceCollection.AddServiceInstance(typeOfInterface, implementation)
	if result != serviceCollection {
		t.Fatal("AddServiceInstance deve retornar a própria instância de serviceCollection")
	}

	// Testando erro ao adicionar implementação inválida
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Adicionar implementação inválida deve causar panic")
		}
	}()
	serviceCollection.AddServiceInstance(typeOfInterface, &MockIncorrectImplementation{})
}

func TestGetServiceByType(t *testing.T) {
	serviceCollection := NewServiceCollection()
	implementation := &MockImplementation{}
	typeOfInterface := reflect.TypeOf((*MockInterface)(nil)).Elem()

	// Adicionar serviço
	serviceCollection.AddServiceInstance(typeOfInterface, implementation)

	// Testar recuperação de serviço
	service := serviceCollection.GetServiceByType(typeOfInterface)
	if service != implementation {
		t.Fatal("GetServiceByType deve retornar a implementação registrada")
	}

	// Testar recuperação de serviço não registrado
	nonExistentType := reflect.TypeOf((*error)(nil)).Elem()
	service = serviceCollection.GetServiceByType(nonExistentType)
	if service != nil {
		t.Fatal("GetServiceByType deve retornar nil para tipos não registrados")
	}
}

func TestAddService(t *testing.T) {
	serviceCollection := NewServiceCollection()
	implementation := &MockImplementation{}

	// Testar função auxiliar genérica AddService
	result := AddService[MockInterface](serviceCollection, implementation)
	if result != serviceCollection {
		t.Fatal("AddService deve retornar a própria instância de serviceCollection")
	}

	// Verificar se o serviço foi adicionado corretamente
	typeOfInterface := reflect.TypeOf((*MockInterface)(nil)).Elem()
	service := serviceCollection.GetServiceByType(typeOfInterface)
	if service != implementation {
		t.Fatal("AddService deve registrar a implementação corretamente")
	}
}

func TestGetService(t *testing.T) {
	serviceCollection := NewServiceCollection()
	implementation := &MockImplementation{}

	// Adicionar serviço
	AddService[MockInterface](serviceCollection, implementation)

	// Testar função auxiliar genérica GetService
	service := GetService[MockInterface](serviceCollection)
	if service == nil {
		t.Fatal("GetService deve retornar a implementação registrada")
	}
	mock, ok := service.(*MockImplementation)
	if !ok || mock != implementation {
		t.Fatal("GetService deve retornar a implementação correta")
	}

	// Testar GetService para tipo não registrado
	var nilService = GetService[error](serviceCollection)
	if nilService != nil {
		t.Fatal("GetService deve retornar zero value para tipos não registrados")
	}
} 