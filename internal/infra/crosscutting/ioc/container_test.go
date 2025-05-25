package ioc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContainer(t *testing.T) {
	// Execução
	container := NewContainer()

	// Verificação
	assert.NotNil(t, container)
	assert.NotNil(t, container.GetUserRepository())
}

func TestContainer_GetUserRepository(t *testing.T) {
	// Configuração
	container := NewContainer()

	// Execução
	repo := container.GetUserRepository()

	// Verificação
	assert.NotNil(t, repo)
}
