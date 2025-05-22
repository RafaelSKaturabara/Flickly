package users

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

// MockMediatorForRouterTest é um mock do mediator para testes do roteador
type MockMediatorForRouterTest struct {
	RegisteredHandlers map[string]mediator.Handler
}

func NewMockMediatorForRouterTest() *MockMediatorForRouterTest {
	return &MockMediatorForRouterTest{
		RegisteredHandlers: make(map[string]mediator.Handler),
	}
}

func (m *MockMediatorForRouterTest) Register(requestName string, handler mediator.Handler) {
	m.RegisteredHandlers[requestName] = handler
}

func (m *MockMediatorForRouterTest) Send(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	if handler, exists := m.RegisteredHandlers[reflect.TypeOf(request).Name()]; exists {
		return handler.Handle(c, request)
	}
	return nil, nil
}

// MockUserRepositoryForRouterTest é um mock do repositório de usuários para testes
type MockUserRepositoryForRouterTest struct{}

func (m *MockUserRepositoryForRouterTest) CreateUser(ctx context.Context, user *entities.User) error {
	return nil
}

func (m *MockUserRepositoryForRouterTest) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserRepositoryForRouterTest) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserRepositoryForRouterTest) UpdateUser(ctx context.Context, user *entities.User) error {
	return nil
}

func (m *MockUserRepositoryForRouterTest) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (m *MockUserRepositoryForRouterTest) GetUserByProviderID(ctx context.Context, provider, providerID string) (*entities.User, error) {
	return nil, nil
}

func (m *MockUserRepositoryForRouterTest) UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error {
	return nil
}

func (m *MockUserRepositoryForRouterTest) UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error {
	return nil
}

func TestStartup(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	router := gin.New()
	serviceCollection := utilities.NewServiceCollection()

	// Registrar as dependências necessárias
	utilities.AddService[mediator.Mediator](serviceCollection, NewMockMediatorForRouterTest())
	utilities.AddService[repositories.IUserRepository](serviceCollection, &MockUserRepositoryForRouterTest{})

	// Execução
	Startup(router, serviceCollection)

	// Verificações
	routes := router.Routes()

	// Verificar se as rotas foram registradas
	var foundPostUser, foundPostOauthToken bool
	for _, route := range routes {
		if route.Path == "/user" && route.Method == "POST" {
			foundPostUser = true
		}
		if route.Path == "/oauth/token" && route.Method == "POST" {
			foundPostOauthToken = true
		}
	}

	assert.True(t, foundPostUser, "A rota POST /user deve estar registrada")
	assert.True(t, foundPostOauthToken, "A rota POST /oauth/token deve estar registrada")
}
