package repositories

import (
	"errors"
	"flickly/internal/domain/users/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockUserRepository é uma implementação mock da interface IUserRepository
type MockUserRepository struct {
	CreateUserCalled     bool
	GetUserByEmailCalled bool
	UserToReturn         *entities.User
	ErrorToReturn        error
}

func (m *MockUserRepository) CreateUser(user *entities.User) error {
	m.CreateUserCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	m.GetUserByEmailCalled = true
	return m.UserToReturn, m.ErrorToReturn
}

func TestIUserRepository_Interface(t *testing.T) {
	// Teste para verificar se a implementação mock satisfaz a interface
	var _ IUserRepository = (*MockUserRepository)(nil)
}

func TestIUserRepository_CreateUser(t *testing.T) {
	// Configuração
	mockRepo := &MockUserRepository{
		ErrorToReturn: nil,
	}
	user := entities.NewUser("Test User", "test@example.com")

	// Execução
	err := mockRepo.CreateUser(user)

	// Verificações
	assert.NoError(t, err, "CreateUser deve retornar nil quando não há erro")
	assert.True(t, mockRepo.CreateUserCalled, "O método CreateUser deve ser chamado")

	// Teste com erro
	expectedError := errors.New("user already exists")
	mockRepo = &MockUserRepository{
		ErrorToReturn: expectedError,
	}

	// Execução
	err = mockRepo.CreateUser(user)

	// Verificações
	assert.Equal(t, expectedError, err, "CreateUser deve retornar o erro esperado")
	assert.True(t, mockRepo.CreateUserCalled, "O método CreateUser deve ser chamado")
}

func TestIUserRepository_GetUserByEmail(t *testing.T) {
	// Configuração
	expectedUser := entities.NewUser("Test User", "test@example.com")
	mockRepo := &MockUserRepository{
		UserToReturn:  expectedUser,
		ErrorToReturn: nil,
	}

	// Execução
	user, err := mockRepo.GetUserByEmail("test@example.com")

	// Verificações
	assert.NoError(t, err, "GetUserByEmail deve retornar nil quando não há erro")
	assert.Equal(t, expectedUser, user, "GetUserByEmail deve retornar o usuário esperado")
	assert.True(t, mockRepo.GetUserByEmailCalled, "O método GetUserByEmail deve ser chamado")

	// Teste com erro
	expectedError := errors.New("database error")
	mockRepo = &MockUserRepository{
		UserToReturn:  nil,
		ErrorToReturn: expectedError,
	}

	// Execução
	user, err = mockRepo.GetUserByEmail("test@example.com")

	// Verificações
	assert.Equal(t, expectedError, err, "GetUserByEmail deve retornar o erro esperado")
	assert.Nil(t, user, "GetUserByEmail deve retornar nil quando há erro")
	assert.True(t, mockRepo.GetUserByEmailCalled, "O método GetUserByEmail deve ser chamado")
} 