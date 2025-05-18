package ioc

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockMediatorForTest é um mock do mediator para testar o injetor de handlers
type MockMediatorForTest struct {
	RegisteredHandlers map[string]mediator.Handler
}

func NewMockMediatorForTest() *MockMediatorForTest {
	return &MockMediatorForTest{
		RegisteredHandlers: make(map[string]mediator.Handler),
	}
}

func (m *MockMediatorForTest) Register(requestName string, handler mediator.Handler) {
	m.RegisteredHandlers[requestName] = handler
}

func (m *MockMediatorForTest) Send(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	return nil, nil
}

// MockUserRepositoryForTest é um mock do repositório de usuários para testes
type MockUserRepositoryForTest struct{}

func (m *MockUserRepositoryForTest) CreateUser(user *entities.User) error {
	return nil
}

func (m *MockUserRepositoryForTest) GetUserByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func TestInjectMediatorHandlers(t *testing.T) {
	// Configuração
	serviceCollection := utilities.NewServiceCollection()
	mockMediator := NewMockMediatorForTest()
	mockUserRepo := &MockUserRepositoryForTest{}

	// Registrar o mediator mock e o repositório necessário para os handlers
	utilities.AddService[mediator.Mediator](serviceCollection, mockMediator)
	utilities.AddService[repositories.IUserRepository](serviceCollection, mockUserRepo)

	// Execução
	InjectMediatorHandlers(serviceCollection)

	// Verificações
	// Verificar se o handler do CreateUserCommand foi registrado
	handler, exists := mockMediator.RegisteredHandlers["CreateUserCommand"]
	assert.True(t, exists, "O handler de CreateUserCommand deve ser registrado")
	assert.NotNil(t, handler, "O handler registrado não deve ser nulo")
}
