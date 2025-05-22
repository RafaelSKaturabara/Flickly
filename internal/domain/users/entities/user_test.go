package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	// Configuração
	name := "Test User"
	email := "test@example.com"
	provider := "google"
	providerID := "123456789"

	// Execução
	user := NewUser(name, email, provider, providerID)

	// Verificações
	assert.NotNil(t, user, "NewUser deve retornar uma instância não nula")
	assert.Equal(t, name, user.Name, "O nome do usuário deve ser configurado corretamente")
	assert.Equal(t, email, user.Email, "O email do usuário deve ser configurado corretamente")
	assert.Equal(t, provider, user.Provider, "O provider deve ser configurado corretamente")
	assert.Equal(t, providerID, user.ProviderID, "O providerID deve ser configurado corretamente")
	assert.False(t, user.VerifiedEmail, "VerifiedEmail deve ser falso para um novo usuário")
	assert.Equal(t, []string{"user"}, user.Roles, "O usuário deve ter o role 'user' por padrão")
	assert.Empty(t, user.Scopes, "O usuário não deve ter escopos inicialmente")

	// Verificar se a entidade base foi inicializada corretamente
	assert.NotEqual(t, uuid.Nil, user.ID, "O ID deve ser inicializado com um UUID válido")
	assert.False(t, user.CreatedAt.IsZero(), "CreatedAt deve ser inicializado com a data atual")
	assert.Nil(t, user.LastUpdateAt, "LastUpdateAt deve ser nulo para um novo usuário")
	assert.Nil(t, user.DeletedAt, "DeletedAt deve ser nulo para um novo usuário")
}

func TestUserRoles(t *testing.T) {
	// Configuração
	user := NewUser("Test User", "test@example.com", "google", "123456789")

	// Teste HasRole
	assert.True(t, user.HasRole("user"), "O usuário deve ter o role 'user' por padrão")
	assert.False(t, user.HasRole("admin"), "O usuário não deve ter o role 'admin' inicialmente")

	// Teste AddRole
	user.AddRole("admin")
	assert.True(t, user.HasRole("admin"), "O usuário deve ter o role 'admin' após adicioná-lo")
	assert.Equal(t, 2, len(user.Roles), "O usuário deve ter exatamente 2 roles")

	// Teste RemoveRole
	user.RemoveRole("user")
	assert.False(t, user.HasRole("user"), "O usuário não deve ter o role 'user' após removê-lo")
	assert.True(t, user.HasRole("admin"), "O usuário deve manter o role 'admin'")
	assert.Equal(t, 1, len(user.Roles), "O usuário deve ter exatamente 1 role")
}

func TestUserScopes(t *testing.T) {
	// Configuração
	user := NewUser("Test User", "test@example.com", "google", "123456789")

	// Teste HasScope
	assert.False(t, user.HasScope("email"), "O usuário não deve ter o escopo 'email' inicialmente")

	// Teste AddScope
	user.AddScope("email")
	assert.True(t, user.HasScope("email"), "O usuário deve ter o escopo 'email' após adicioná-lo")
	assert.Equal(t, 1, len(user.Scopes), "O usuário deve ter exatamente 1 escopo")

	// Teste RemoveScope
	user.RemoveScope("email")
	assert.False(t, user.HasScope("email"), "O usuário não deve ter o escopo 'email' após removê-lo")
	assert.Equal(t, 0, len(user.Scopes), "O usuário não deve ter escopos")
}

func TestUpdateOAuthInfo(t *testing.T) {
	// Configuração
	user := NewUser("Test User", "test@example.com", "google", "123456789")
	accessToken := "access_token_123"
	refreshToken := "refresh_token_123"
	tokenExpiry := int64(1234567890)
	scopes := []string{"email", "profile"}

	// Execução
	user.UpdateOAuthInfo(accessToken, refreshToken, tokenExpiry, scopes)

	// Verificações
	assert.Equal(t, accessToken, user.AccessToken, "AccessToken deve ser atualizado corretamente")
	assert.Equal(t, refreshToken, user.RefreshToken, "RefreshToken deve ser atualizado corretamente")
	assert.Equal(t, tokenExpiry, user.TokenExpiry, "TokenExpiry deve ser atualizado corretamente")
	assert.Equal(t, scopes, user.Scopes, "Scopes deve ser atualizado corretamente")
}

func TestUpdateProfile(t *testing.T) {
	// Configuração
	user := NewUser("Test User", "test@example.com", "google", "123456789")
	newName := "Updated Name"
	givenName := "Updated"
	familyName := "Name"
	picture := "https://example.com/picture.jpg"

	// Execução
	user.UpdateProfile(newName, givenName, familyName, picture, true)

	// Verificações
	assert.Equal(t, newName, user.Name, "Name deve ser atualizado corretamente")
	assert.Equal(t, givenName, user.GivenName, "GivenName deve ser atualizado corretamente")
	assert.Equal(t, familyName, user.FamilyName, "FamilyName deve ser atualizado corretamente")
	assert.Equal(t, picture, user.Picture, "Picture deve ser atualizado corretamente")
	assert.True(t, user.VerifiedEmail, "VerifiedEmail deve ser atualizado corretamente")
}
