package ioc

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInjectServices(t *testing.T) {
	// Configuração
	serviceCollection := utilities.NewServiceCollection()

	// Execução
	InjectServices(serviceCollection)

	// Verificações
	// Verificar se o mediator foi registrado
	mediatR := utilities.GetService[mediator.Mediator](serviceCollection)
	assert.NotNil(t, mediatR, "O mediator deve ser registrado")

	// Verificar se o repositório de usuários foi registrado
	userRepo := utilities.GetService[repositories.IUserRepository](serviceCollection)
	assert.NotNil(t, userRepo, "O repositório de usuários deve ser registrado")
}
