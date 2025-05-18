package entities

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	// Configuração
	name := "Test User"
	email := "test@example.com"

	// Execução
	user := NewUser(name, email)

	// Verificações
	assert.NotNil(t, user, "NewUser deve retornar uma instância não nula")
	assert.Equal(t, name, user.Name, "O nome do usuário deve ser configurado corretamente")
	assert.Equal(t, email, user.Email, "O email do usuário deve ser configurado corretamente")
	
	// Verificar se a entidade base foi inicializada corretamente
	assert.NotEqual(t, uuid.Nil, user.ID, "O ID deve ser inicializado com um UUID válido")
	assert.False(t, user.CreatedAt.IsZero(), "CreatedAt deve ser inicializado com a data atual")
	assert.Nil(t, user.LastUpdateAt, "LastUpdateAt deve ser nulo para um novo usuário")
	assert.Nil(t, user.DeletedAt, "DeletedAt deve ser nulo para um novo usuário")
} 