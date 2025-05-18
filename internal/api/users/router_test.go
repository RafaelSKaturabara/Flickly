package users

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

// MockMediatorForRouterTest é um mock do mediator para testes do roteador
type MockMediatorForRouterTest struct{}

func (m *MockMediatorForRouterTest) Register(requestName string, handler mediator.Handler) {}

func (m *MockMediatorForRouterTest) Send(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	return nil, nil
}

// MockUserRepositoryForRouterTest é um mock do repositório de usuários para testes
type MockUserRepositoryForRouterTest struct{}

func (m *MockUserRepositoryForRouterTest) CreateUser(user *entities.User) error {
	return nil
}

func (m *MockUserRepositoryForRouterTest) GetUserByEmail(email string) (*entities.User, error) {
	return nil, nil
}

func TestStartup(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	router := gin.New()
	serviceCollection := utilities.NewServiceCollection()

	// Registrar as dependências necessárias
	utilities.AddService[mediator.Mediator](serviceCollection, &MockMediatorForRouterTest{})
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
