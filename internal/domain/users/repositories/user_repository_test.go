package repositories

import (
	"context"
	"errors"
	"flickly/internal/domain/users/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// MockUserRepository é uma implementação mock da interface IUserRepository
type MockUserRepository struct {
	CreateUserCalled          bool
	GetUserByEmailCalled      bool
	GetUserByIDCalled         bool
	UpdateUserCalled          bool
	DeleteUserCalled          bool
	GetUserByProviderIDCalled bool
	UpdateUserOAuthInfoCalled bool
	UpdateUserRolesCalled     bool
	UserToReturn              *entities.User
	ErrorToReturn             error
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	m.CreateUserCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	m.GetUserByEmailCalled = true
	return m.UserToReturn, m.ErrorToReturn
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	m.GetUserByIDCalled = true
	return m.UserToReturn, m.ErrorToReturn
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	m.UpdateUserCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	m.DeleteUserCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepository) GetUserByProviderID(ctx context.Context, provider, providerID string) (*entities.User, error) {
	m.GetUserByProviderIDCalled = true
	return m.UserToReturn, m.ErrorToReturn
}

func (m *MockUserRepository) UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error {
	m.UpdateUserOAuthInfoCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepository) UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error {
	m.UpdateUserRolesCalled = true
	return m.ErrorToReturn
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
	user := entities.NewUser("Test User", "test@example.com", "google", "123456789")
	ctx := context.Background()

	// Execução
	err := mockRepo.CreateUser(ctx, user)

	// Verificações
	assert.NoError(t, err, "CreateUser deve retornar nil quando não há erro")
	assert.True(t, mockRepo.CreateUserCalled, "O método CreateUser deve ser chamado")

	// Teste com erro
	expectedError := errors.New("user already exists")
	mockRepo = &MockUserRepository{
		ErrorToReturn: expectedError,
	}

	// Execução
	err = mockRepo.CreateUser(ctx, user)

	// Verificações
	assert.Equal(t, expectedError, err, "CreateUser deve retornar o erro esperado")
	assert.True(t, mockRepo.CreateUserCalled, "O método CreateUser deve ser chamado")
}

func TestIUserRepository_GetUserByEmail(t *testing.T) {
	// Configuração
	expectedUser := entities.NewUser("Test User", "test@example.com", "google", "123456789")
	mockRepo := &MockUserRepository{
		UserToReturn:  expectedUser,
		ErrorToReturn: nil,
	}
	ctx := context.Background()

	// Execução
	user, err := mockRepo.GetUserByEmail(ctx, "test@example.com")

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
	user, err = mockRepo.GetUserByEmail(ctx, "test@example.com")

	// Verificações
	assert.Equal(t, expectedError, err, "GetUserByEmail deve retornar o erro esperado")
	assert.Nil(t, user, "GetUserByEmail deve retornar nil quando há erro")
	assert.True(t, mockRepo.GetUserByEmailCalled, "O método GetUserByEmail deve ser chamado")
}

func TestIUserRepository_GetUserByProviderID(t *testing.T) {
	// Configuração
	expectedUser := entities.NewUser("Test User", "test@example.com", "google", "123456789")
	mockRepo := &MockUserRepository{
		UserToReturn:  expectedUser,
		ErrorToReturn: nil,
	}
	ctx := context.Background()

	// Execução
	user, err := mockRepo.GetUserByProviderID(ctx, "google", "123456789")

	// Verificações
	assert.NoError(t, err, "GetUserByProviderID deve retornar nil quando não há erro")
	assert.Equal(t, expectedUser, user, "GetUserByProviderID deve retornar o usuário esperado")
	assert.True(t, mockRepo.GetUserByProviderIDCalled, "O método GetUserByProviderID deve ser chamado")

	// Teste com erro
	expectedError := errors.New("user not found")
	mockRepo = &MockUserRepository{
		UserToReturn:  nil,
		ErrorToReturn: expectedError,
	}

	// Execução
	user, err = mockRepo.GetUserByProviderID(ctx, "google", "123456789")

	// Verificações
	assert.Equal(t, expectedError, err, "GetUserByProviderID deve retornar o erro esperado")
	assert.Nil(t, user, "GetUserByProviderID deve retornar nil quando há erro")
	assert.True(t, mockRepo.GetUserByProviderIDCalled, "O método GetUserByProviderID deve ser chamado")
}

func TestIUserRepository_UpdateUserOAuthInfo(t *testing.T) {
	// Configuração
	mockRepo := &MockUserRepository{
		ErrorToReturn: nil,
	}
	ctx := context.Background()

	// Execução
	err := mockRepo.UpdateUserOAuthInfo(ctx, uuid.New(), "access_token", "refresh_token", 1234567890, []string{"email", "profile"})

	// Verificações
	assert.NoError(t, err, "UpdateUserOAuthInfo deve retornar nil quando não há erro")
	assert.True(t, mockRepo.UpdateUserOAuthInfoCalled, "O método UpdateUserOAuthInfo deve ser chamado")

	// Teste com erro
	expectedError := errors.New("update failed")
	mockRepo = &MockUserRepository{
		ErrorToReturn: expectedError,
	}

	// Execução
	err = mockRepo.UpdateUserOAuthInfo(ctx, uuid.New(), "access_token", "refresh_token", 1234567890, []string{"email", "profile"})

	// Verificações
	assert.Equal(t, expectedError, err, "UpdateUserOAuthInfo deve retornar o erro esperado")
	assert.True(t, mockRepo.UpdateUserOAuthInfoCalled, "O método UpdateUserOAuthInfo deve ser chamado")
}

func TestIUserRepository_UpdateUserRoles(t *testing.T) {
	// Configuração
	mockRepo := &MockUserRepository{
		ErrorToReturn: nil,
	}
	ctx := context.Background()

	// Execução
	err := mockRepo.UpdateUserRoles(ctx, uuid.New(), []string{"admin", "user"})

	// Verificações
	assert.NoError(t, err, "UpdateUserRoles deve retornar nil quando não há erro")
	assert.True(t, mockRepo.UpdateUserRolesCalled, "O método UpdateUserRoles deve ser chamado")

	// Teste com erro
	expectedError := errors.New("update failed")
	mockRepo = &MockUserRepository{
		ErrorToReturn: expectedError,
	}

	// Execução
	err = mockRepo.UpdateUserRoles(ctx, uuid.New(), []string{"admin", "user"})

	// Verificações
	assert.Equal(t, expectedError, err, "UpdateUserRoles deve retornar o erro esperado")
	assert.True(t, mockRepo.UpdateUserRolesCalled, "O método UpdateUserRoles deve ser chamado")
}
