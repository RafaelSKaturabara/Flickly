package integration_tests

import (
	"flickly/internal/domain/core/mediator"
	"flickly/internal/domain/users/entities"
	"flickly/internal/domain/users/repositories"
	inversionofcontrol "flickly/internal/infra/cross-cutting/inversion-of-control"
	"flickly/internal/infra/cross-cutting/utilities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// DataIntegrationTestSuite define a suite de testes para a camada de dados
type DataIntegrationTestSuite struct {
	suite.Suite
	serviceCollection utilities.IServiceCollection
	userRepository    repositories.IUserRepository
}

// SetupSuite configura a suite de testes
func (suite *DataIntegrationTestSuite) SetupSuite() {
	// Inicializar o service collection
	serviceCollection := utilities.NewServiceCollection()
	
	// Injetar serviços reais (não mocks)
	inversionofcontrol.InjectServices(serviceCollection)
	
	// Obter o repositório de usuários
	userRepository := utilities.GetService[repositories.IUserRepository](serviceCollection)
	
	suite.serviceCollection = serviceCollection
	suite.userRepository = userRepository
}

// TestUserRepositoryOperations testa as operações do repositório de usuários
func (suite *DataIntegrationTestSuite) TestUserRepositoryOperations() {
	// Criar um usuário de teste
	email := "integracao@example.com"
	user := entities.NewUser("Usuário Integração", email)
	
	// Salvar o usuário
	err := suite.userRepository.CreateUser(user)
	assert.NoError(suite.T(), err)
	
	// Recuperar o usuário pelo email
	retrievedUser, err := suite.userRepository.GetUserByEmail(email)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedUser)
	assert.Equal(suite.T(), user.Name, retrievedUser.Name)
	assert.Equal(suite.T(), user.Email, retrievedUser.Email)
}

// TestMediatorIntegration testa a integração do mediator com os handlers
func (suite *DataIntegrationTestSuite) TestMediatorIntegration() {
	// Este teste verifica a integração entre o mediator e os handlers registrados
	// Esta é uma parte crucial da arquitetura do aplicativo
	
	// Verificar se o mediator foi configurado corretamente
	mediator := utilities.GetService[mediator.Mediator](suite.serviceCollection)
	assert.NotNil(suite.T(), mediator)
	
	// Os detalhes específicos deste teste dependerão da implementação do mediator
	// e dos handlers registrados no seu aplicativo
}

// TestRunDataSuite executa a suite de testes
func TestRunDataSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Pulando testes de integração em modo curto")
	}
	suite.Run(t, new(DataIntegrationTestSuite))
} 