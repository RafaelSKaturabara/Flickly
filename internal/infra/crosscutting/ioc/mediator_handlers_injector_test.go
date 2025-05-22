package ioc

import (
	"context"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
	if handler, exists := m.RegisteredHandlers[reflect.TypeOf(request).Name()]; exists {
		return handler.Handle(c, request)
	}
	return nil, nil
}

// MockUserRepositoryForTest é um mock do repositório de usuários para testes
type MockUserRepositoryForTest struct{}

func (m *MockUserRepositoryForTest) CreateUser(ctx context.Context, user *entities.User) error {
	return nil
}

func (m *MockUserRepositoryForTest) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserRepositoryForTest) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserRepositoryForTest) UpdateUser(ctx context.Context, user *entities.User) error {
	return nil
}

func (m *MockUserRepositoryForTest) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (m *MockUserRepositoryForTest) GetUserByProviderID(ctx context.Context, provider, providerID string) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserRepositoryForTest) UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error {
	return nil
}

func (m *MockUserRepositoryForTest) UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error {
	return nil
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
