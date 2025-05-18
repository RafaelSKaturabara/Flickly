package commands

import (
	"errors"
	"flickly/internal/domain/core"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockUserRepository é um mock do repositório de usuários para os testes
type MockUserRepository struct {
	CreateUserCalled bool
	UserToReturn     *entities.User
	ErrorToReturn    error
}

func (m *MockUserRepository) CreateUser(user *entities.User) error {
	m.CreateUserCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	return m.UserToReturn, m.ErrorToReturn
}

// MockMediator é um mock do mediator para os testes
type MockMediator struct {
	RegisterCalled   bool
	SendCalled       bool
	ResponseToReturn mediator.Response
	ErrorToReturn    error
}

func (m *MockMediator) Register(requestName string, handler mediator.Handler) {
	m.RegisterCalled = true
}

func (m *MockMediator) Send(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	m.SendCalled = true
	return m.ResponseToReturn, m.ErrorToReturn
}

// Criando um ServiceCollection com mocks para os testes
func setupMockServices(mockRepo *MockUserRepository, mockMediator *MockMediator) utilities.IServiceCollection {
	serviceCollection := utilities.NewServiceCollection()

	// Registrar o mock do repositório
	utilities.AddService[repositories.IUserRepository](serviceCollection, mockRepo)

	// Registrar o mock do mediator
	utilities.AddService[mediator.Mediator](serviceCollection, mockMediator)

	return serviceCollection
}

func TestNewCreateUserCommandHandler(t *testing.T) {
	// Configuração
	mockRepo := &MockUserRepository{}
	mockMediator := &MockMediator{}
	serviceCollection := setupMockServices(mockRepo, mockMediator)

	// Execução
	handler := NewCreateUserCommandHandler(serviceCollection)

	// Verificações
	assert.NotNil(t, handler, "NewCreateUserCommandHandler deve retornar uma instância não nula")
	assert.NotNil(t, handler.userRepository, "O repositório no handler deve ser inicializado")
	assert.NotNil(t, handler.mediator, "O mediator no handler deve ser inicializado")
}

func TestHandle_Success(t *testing.T) {
	// Configuração
	mockRepo := &MockUserRepository{
		ErrorToReturn: nil,
	}
	mockMediator := &MockMediator{}
	serviceCollection := setupMockServices(mockRepo, mockMediator)

	handler := NewCreateUserCommandHandler(serviceCollection)
	command := CreateUserCommand{
		Name:  "Test User",
		Email: "test@example.com",
	}

	// Execução
	ginContext, _ := gin.CreateTestContext(nil)
	response, err := handler.Handle(ginContext, command)

	// Verificações
	assert.NoError(t, err, "Handle não deve retornar erro quando o repositório retorna sucesso")
	assert.NotNil(t, response, "Response não deve ser nil em caso de sucesso")

	user, ok := response.(*entities.User)
	assert.True(t, ok, "Response deve ser do tipo *entities.User")
	assert.Equal(t, command.Name, user.Name, "O nome do usuário na resposta deve corresponder ao comando")
	assert.Equal(t, command.Email, user.Email, "O email do usuário na resposta deve corresponder ao comando")

	assert.True(t, mockRepo.CreateUserCalled, "O método CreateUser do repositório deve ser chamado")
}

func TestHandle_Error(t *testing.T) {
	// Configuração
	mockError := errors.New("user already exists")
	mockRepo := &MockUserRepository{
		ErrorToReturn: mockError,
	}
	mockMediator := &MockMediator{}
	serviceCollection := setupMockServices(mockRepo, mockMediator)

	handler := NewCreateUserCommandHandler(serviceCollection)
	command := CreateUserCommand{
		Name:  "Test User",
		Email: "existing@example.com",
	}

	// Execução
	ginContext, _ := gin.CreateTestContext(nil)
	response, err := handler.Handle(ginContext, command)

	// Verificações
	assert.Error(t, err, "Handle deve retornar erro quando o repositório retorna erro")
	assert.Nil(t, response, "Response deve ser nil em caso de erro")

	// Verificar se o erro retornado é um DomainError
	domainErr, ok := err.(*core.DomainError)
	assert.True(t, ok, "Erro retornado deve ser do tipo *core.DomainError")
	assert.Equal(t, "Usuário já cadastrado", domainErr.Message, "A mensagem de erro deve ser correta")
	assert.Equal(t, 1, domainErr.Code, "O código de erro deve ser 1")

	assert.True(t, mockRepo.CreateUserCalled, "O método CreateUser do repositório deve ser chamado")
}
