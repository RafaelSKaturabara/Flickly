package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	viewmodels "flickly/internal/api/users/viewmodels"
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/commands"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// MockMediatorForControllerTest é um mock do mediator para testes do controlador
type MockMediatorForControllerTest struct {
	SendCalled       bool
	ResponseToReturn mediator.Response
	ErrorToReturn    error
}

func (m *MockMediatorForControllerTest) Register(requestName string, handler mediator.Handler) {
}

func (m *MockMediatorForControllerTest) Send(c *gin.Context, request mediator.Request) (mediator.Response, error) {
	m.SendCalled = true
	return m.ResponseToReturn, m.ErrorToReturn
}

// MockUserRepositoryForControllerTest é um mock do repositório de usuários para testes
type MockUserRepositoryForControllerTest struct {
	GetUserByEmailCalled bool
	CreateUserCalled     bool
	UserToReturn         *entities.User
	ErrorToReturn        error
}

func (m *MockUserRepositoryForControllerTest) CreateUser(user *entities.User) error {
	m.CreateUserCalled = true
	return m.ErrorToReturn
}

func (m *MockUserRepositoryForControllerTest) GetUserByEmail(email string) (*entities.User, error) {
	m.GetUserByEmailCalled = true
	return m.UserToReturn, m.ErrorToReturn
}

// Função para configurar as dependências de teste
func setupTestDependencies(
	mockMediator *MockMediatorForControllerTest,
	mockRepo *MockUserRepositoryForControllerTest,
) utilities.IServiceCollection {
	serviceCollection := utilities.NewServiceCollection()
	utilities.AddService[mediator.Mediator](serviceCollection, mockMediator)
	utilities.AddService[repositories.IUserRepository](serviceCollection, mockRepo)
	return serviceCollection
}

func TestNewUserController(t *testing.T) {
	// Configuração
	mockMediator := &MockMediatorForControllerTest{}
	mockRepo := &MockUserRepositoryForControllerTest{}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	// Execução
	controller := NewUserController(serviceCollection)

	// Verificações
	assert.NotNil(t, controller, "NewUserController deve retornar uma instância não nula")
	assert.NotNil(t, controller.mediator, "O mediator no controller deve ser inicializado")
	assert.NotNil(t, controller.userRepository, "O repositório no controller deve ser inicializado")
}

func TestPostUser_Success(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	mockMediator := &MockMediatorForControllerTest{
		ResponseToReturn: entities.NewUser("Test User", "test@example.com"),
	}
	mockRepo := &MockUserRepositoryForControllerTest{}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	controller := NewUserController(serviceCollection)

	// Criar request e response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Criar corpo da requisição
	createUserCommand := commands.CreateUserCommand{
		Name:  "Test User",
		Email: "test@example.com",
	}
	jsonData, _ := json.Marshal(createUserCommand)
	c.Request = httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	// Execução
	controller.PostUser(c)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code, "O código de status deve ser 200 OK")
	assert.True(t, mockMediator.SendCalled, "O método Send do mediator deve ser chamado")

	// Verificar o corpo da resposta
	var response entities.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "Test User", response.Name, "O nome do usuário na resposta deve ser correto")
	assert.Equal(t, "test@example.com", response.Email, "O email do usuário na resposta deve ser correto")
}

func TestPostUser_BindError(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	mockMediator := &MockMediatorForControllerTest{}
	mockRepo := &MockUserRepositoryForControllerTest{}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	controller := NewUserController(serviceCollection)

	// Criar request e response recorder com JSON inválido
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/user", strings.NewReader("{invalid json}"))
	c.Request.Header.Set("Content-Type", "application/json")

	// Execução
	controller.PostUser(c)

	// Verificações
	assert.Equal(t, http.StatusBadRequest, w.Code, "O código de status deve ser 400 Bad Request")
	assert.False(t, mockMediator.SendCalled, "O método Send do mediator não deve ser chamado")
}

func TestPostUser_MediatorError(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	expectedError := errors.New("mediator error")
	mockMediator := &MockMediatorForControllerTest{
		ErrorToReturn: expectedError,
	}
	mockRepo := &MockUserRepositoryForControllerTest{}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	controller := NewUserController(serviceCollection)

	// Criar request e response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Criar corpo da requisição
	createUserCommand := commands.CreateUserCommand{
		Name:  "Test User",
		Email: "test@example.com",
	}
	jsonData, _ := json.Marshal(createUserCommand)
	c.Request = httptest.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	// Execução
	controller.PostUser(c)

	// Verificações
	assert.Equal(t, http.StatusBadRequest, w.Code, "O código de status deve ser 400 Bad Request")
	assert.True(t, mockMediator.SendCalled, "O método Send do mediator deve ser chamado")
}

func TestPostOauthToken_Success(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	mockMediator := &MockMediatorForControllerTest{}
	mockRepo := &MockUserRepositoryForControllerTest{
		UserToReturn: entities.NewUser("Test User", "test@example.com"),
	}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	controller := NewUserController(serviceCollection)

	// Criar request e response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Criar dados de formulário para a requisição OAuth
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("client_id", "my_client_id")
	form.Add("client_secret", "my_client_secret")
	form.Add("username", "test@example.com")
	form.Add("password", "password123")

	c.Request = httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execução
	controller.PostOauthToken(c)

	// Verificações
	assert.Equal(t, http.StatusOK, w.Code, "O código de status deve ser 200 OK")
	assert.True(t, mockRepo.GetUserByEmailCalled, "O método GetUserByEmail do repositório deve ser chamado")

	// Verificar o corpo da resposta
	var response viewmodels.TokenResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "some_generated_token", response.AccessToken, "O token de acesso deve ser correto")
	assert.Equal(t, "Bearer Test User", response.TokenType, "O tipo de token deve ser correto")
	assert.Equal(t, 3600, response.ExpiresIn, "O tempo de expiração deve ser correto")
}

func TestPostOauthToken_InvalidCredentials(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	mockMediator := &MockMediatorForControllerTest{}
	mockRepo := &MockUserRepositoryForControllerTest{
		UserToReturn: nil, // Usuário não encontrado
	}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	controller := NewUserController(serviceCollection)

	// Criar request e response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Criar dados de formulário para a requisição OAuth com credenciais inválidas
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("client_id", "wrong_client_id") // ID de cliente incorreto
	form.Add("client_secret", "my_client_secret")
	form.Add("username", "test@example.com")
	form.Add("password", "password123")

	c.Request = httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execução
	controller.PostOauthToken(c)

	// Verificações
	assert.Equal(t, http.StatusUnauthorized, w.Code, "O código de status deve ser 401 Unauthorized")

	// Verificar o corpo da resposta
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err, "Não deve ocorrer erro ao desserializar a resposta JSON")
	assert.Equal(t, "invalid_grant", response["error"], "A mensagem de erro deve ser 'invalid_grant'")
}

func TestPostOauthToken_RepositoryError(t *testing.T) {
	// Configuração
	gin.SetMode(gin.TestMode)
	expectedError := errors.New("repository error")
	mockMediator := &MockMediatorForControllerTest{}
	mockRepo := &MockUserRepositoryForControllerTest{
		ErrorToReturn: expectedError,
	}
	serviceCollection := setupTestDependencies(mockMediator, mockRepo)

	controller := NewUserController(serviceCollection)

	// Criar request e response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Criar dados de formulário para a requisição OAuth
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("client_id", "my_client_id")
	form.Add("client_secret", "my_client_secret")
	form.Add("username", "test@example.com")
	form.Add("password", "password123")

	c.Request = httptest.NewRequest(http.MethodPost, "/oauth/token", strings.NewReader(form.Encode()))
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Execução
	controller.PostOauthToken(c)

	// Verificações
	assert.Equal(t, http.StatusUnauthorized, w.Code, "O código de status deve ser 401 Unauthorized")
	assert.True(t, mockRepo.GetUserByEmailCalled, "O método GetUserByEmail do repositório deve ser chamado")
}
