package core

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEntity(t *testing.T) {
	// Execução
	entity := NewEntity()

	// Verificações
	assert.NotEqual(t, uuid.Nil, entity.ID, "O ID deve ser inicializado com um UUID válido")
	assert.False(t, entity.CreatedAt.IsZero(), "CreatedAt deve ser inicializado com a data atual")
	assert.Nil(t, entity.LastUpdateAt, "LastUpdateAt deve ser nulo para uma nova entidade")
	assert.Nil(t, entity.DeletedAt, "DeletedAt deve ser nulo para uma nova entidade")
} 