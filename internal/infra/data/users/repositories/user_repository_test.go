package repositories

import (
	"context"
	"testing"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/users/gormmappings"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	_ = db.AutoMigrate(&gormmappings.UserDB{})
	return db
}

func TestNewUserRepository(t *testing.T) {
	// Configuração
	db := setupTestDB()

	// Execução
	repository := NewUserRepository(db)

	// Verificações
	assert.NotNil(t, repository, "NewUserRepository deve retornar uma instância não nula")
}

func TestCreateUser(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	user := entities.NewUser("Test User", "test@example.com", "google", "123456789", "password123")
	ctx := context.Background()

	// Execução - primeiro usuário
	err := repository.Create(user)

	// Verificações
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o primeiro usuário")

	// Verifica se o usuário foi criado buscando-o pelo email
	retrievedUser, err := repository.GetUserByEmail(ctx, user.Email)
	assert.NoError(t, err, "Não deve ocorrer erro ao buscar o usuário criado")
	assert.Equal(t, user.Email, retrievedUser.Email, "O email do usuário deve ser armazenado corretamente")

	// Execução - tentativa de duplicar usuário
	duplicateUser := entities.NewUser("Duplicate User", "test@example.com", "google", "987654321", "password456")
	err = repository.Create(duplicateUser)

	// Verificações
	assert.Error(t, err, "Deve ocorrer erro ao criar usuário com email duplicado")
}

func TestUpdateUser(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	user := entities.NewUser("Test User", "test@example.com", "google", "123456789", "password123")
	err := repository.Create(user)
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o usuário para teste")

	// Execução
	user.Name = "Updated Name"
	err = repository.Update(user)

	// Verificações
	assert.NoError(t, err)
	retrieved, _ := repository.GetByID(user.ID)
	assert.Equal(t, "Updated Name", retrieved.Name)
}

func TestGetUserByEmail(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	user := entities.NewUser("Test User", "test@example.com", "google", "123456789", "password123")
	ctx := context.Background()
	err := repository.Create(user)
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

func TestGetUserByEmailAndPasswordAndClientAndSecret(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	password := "password123"
	hashedPassword, _ := utilities.Encrypt(password)

	user := entities.NewUser("Test User", "test@example.com", "client1", "secret1", hashedPassword)
	ctx := context.Background()
	err := repository.Create(user)
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o usuário para teste")

	// Execução - Sucesso
	retrieved, err := repository.GetUserByEmailAndPasswordAndClientAndSecret(ctx, "test@example.com", password, "client1", "secret1")
	assert.NoError(t, err)
	assert.NotNil(t, retrieved)

	// Execução - Senha errada
	retrieved, err = repository.GetUserByEmailAndPasswordAndClientAndSecret(ctx, "test@example.com", "wrong", "client1", "secret1")
	assert.Error(t, err)
	assert.Nil(t, retrieved)

	// Execução - ClientID errado
	retrieved, err = repository.GetUserByEmailAndPasswordAndClientAndSecret(ctx, "test@example.com", password, "wrong", "secret1")
	assert.Error(t, err)
	assert.Nil(t, retrieved)
}

func TestUpdateUserOAuthInfo(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	user := entities.NewUser("Test User", "test@example.com", "c", "s", "p")
	ctx := context.Background()
	err := repository.Create(user)
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o usuário para teste")

	// Execução
	err = repository.UpdateUserOAuthInfo(ctx, user.ID, "new-token", "refresh", 3600, []string{"scope1"})

	// Verificações
	assert.NoError(t, err)
	retrieved, _ := repository.GetByID(user.ID)
	assert.NotNil(t, retrieved)
	assert.Equal(t, "new-token", retrieved.AccessToken)
	assert.Equal(t, int64(3600), retrieved.TokenExpiry)
	assert.Contains(t, retrieved.TokenScopes, "scope1")
}

func TestUpdateUserRoles(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	user := entities.NewUser("Test User", "test@example.com", "c", "s", "p")
	ctx := context.Background()
	err := repository.Create(user)
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o usuário para teste")

	// Execução
	err = repository.UpdateUserRoles(ctx, user.ID, []string{"admin", "editor"})

	// Verificações
	assert.NoError(t, err)
	retrieved, _ := repository.GetByID(user.ID)
	assert.NotNil(t, retrieved)
	assert.Contains(t, retrieved.Roles, "admin")
	assert.Contains(t, retrieved.Roles, "editor")
}

func TestDeleteUser(t *testing.T) {
	// Configuração
	db := setupTestDB()
	repository := NewUserRepository(db)
	user := entities.NewUser("Test User", "test@example.com", "c", "s", "p")
	err := repository.Create(user)
	assert.NoError(t, err, "Não deve ocorrer erro ao criar o usuário para teste")

	// Execução
	err = repository.Delete(user.ID)

	// Verificações
	assert.NoError(t, err)
	retrieved, _ := repository.GetByID(user.ID)
	assert.Nil(t, retrieved)
}
