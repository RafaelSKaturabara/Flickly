package utilities

import "reflect"

// IServiceCollection interface principal sem genéricos
type IServiceCollection interface {
	AddServiceInstance(serviceType reflect.Type, implementation any) IServiceCollection
	GetServiceByType(serviceType reflect.Type) any
}

type serviceCollection struct {
	services map[reflect.Type]any
}

func NewServiceCollection() IServiceCollection {
	return &serviceCollection{
		services: make(map[reflect.Type]interface{}),
	}
}

// AddServiceInstance implementa a interface não-genérica
func (c *serviceCollection) AddServiceInstance(serviceType reflect.Type, implementation any) IServiceCollection {
	implementationValue := reflect.ValueOf(implementation)

	if !implementationValue.Type().Implements(serviceType) {
		panic("a implementação não satisfaz a interface esperada")
	}

	c.services[serviceType] = implementation
	return c
}

// GetServiceByType implementa a interface não-genérica
func (c *serviceCollection) GetServiceByType(serviceType reflect.Type) any {
	if service, ok := c.services[serviceType]; ok {
		return service
	}
	return nil
}

// Funções auxiliares genéricas (não são parte da interface)

// AddService é uma função helper genérica para adicionar serviços
func AddService[T any](container IServiceCollection, implementation interface{}) IServiceCollection {
	typeOfInterface := reflect.TypeOf((*T)(nil)).Elem()
	return container.AddServiceInstance(typeOfInterface, implementation)
}

// GetService é uma função helper genérica para obter serviços
func GetService[T any](container IServiceCollection) T {
	typeOfInterface := reflect.TypeOf((*T)(nil)).Elem()
	service := container.GetServiceByType(typeOfInterface)
	if service == nil {
		var zero T
		return zero
	}
	return service.(T)
}
