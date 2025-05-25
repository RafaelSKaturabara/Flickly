package repositories

import (
	"context"
	"testing"

	"github.com/rkaturabara/flickly/internal/domain/users/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewUserRepository(t *testing.T) {
	// Execução
	repository := NewUserRepository()

	// Verificações
	assert.NotNil(t, repository, "NewUserRepository deve retornar uma instância não nula")
}

func TestCreateUser(t *testing.T) {
	// Configuração
	repository := NewUserRepository()
	user := entities.NewUser("Test User", "test@example.com", "google", "123456789")
	ctx := context.Background()

	// Execução - primeiro usuário
	err := repository.CreateUser(ctx, user)

	// Verificações
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o primeiro usuário")

	// Verifica se o usuário foi criado buscando-o pelo email
	retrievedUser, err := repository.GetUserByEmail(ctx, user.Email)
	assert.NoError(t, err, "Não deve ocorrer erro ao buscar o usuário criado")
	assert.Equal(t, user.Email, retrievedUser.Email, "O email do usuário deve ser armazenado corretamente")

	// Execução - tentativa de duplicar usuário
	duplicateUser := entities.NewUser("Duplicate User", "test@example.com", "google", "987654321")
	err = repository.CreateUser(ctx, duplicateUser)

	// Verificações
	assert.Error(t, err, "Deve ocorrer erro ao criar usuário com email duplicado")
	assert.Equal(t, "user already exists", err.Error(), "Mensagem de erro incorreta")
}

func TestGetUserByEmail(t *testing.T) {
	// Configuração
	repository := NewUserRepository()
	user := entities.NewUser("Test User", "test@example.com", "google", "123456789")
	ctx := context.Background()
	err := repository.CreateUser(ctx, user)
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o usuário para teste")

	// Execução - usuário existente
	retrievedUser, err := repository.GetUserByEmail(ctx, "test@example.com")

	// Verificações
	assert.NoError(t, err, "Não deve ocorrer erro ao buscar usuário existente")
	assert.NotNil(t, retrievedUser, "Deve retornar o usuário quando encontrado")
	assert.Equal(t, user.Email, retrievedUser.Email, "O email do usuário recuperado deve ser igual ao esperado")
	assert.Equal(t, user.Name, retrievedUser.Name, "O nome do usuário recuperado deve ser igual ao esperado")

	// Execução - usuário não existente
	retrievedUser, err = repository.GetUserByEmail(ctx, "nonexistent@example.com")

	// Verificações
	assert.Error(t, err, "Deve ocorrer erro ao buscar usuário não existente")
	assert.Nil(t, retrievedUser, "Deve retornar nil para usuário não encontrado")
}
