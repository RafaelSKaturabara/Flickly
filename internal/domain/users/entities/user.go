package entities

import (
	"flickly/internal/domain/core"
)

// User representa um usuário no sistema com suporte a OAuth2
type User struct {
	core.Entity
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	GivenName     string   `json:"givenName,omitempty"`
	FamilyName    string   `json:"familyName,omitempty"`
	Picture       string   `json:"picture,omitempty"`
	VerifiedEmail bool     `json:"verifiedEmail"`
	Provider      string   `json:"provider"`   // Ex: "google", "github", etc
	ProviderID    string   `json:"providerId"` // ID do usuário no provedor OAuth
	Roles         []string `json:"roles"`
	Scopes        []string `json:"scopes,omitempty"` // Escopos OAuth2 concedidos
	AccessToken   string   `json:"accessToken,omitempty"`
	RefreshToken  string   `json:"refreshToken,omitempty"`
	TokenExpiry   int64    `json:"tokenExpiry,omitempty"`
}

// NewUser cria uma nova instância de User
func NewUser(name, email, provider, providerID string) *User {
	return &User{
		Entity:        core.NewEntity(),
		Name:          name,
		Email:         email,
		Provider:      provider,
		ProviderID:    providerID,
		Roles:         []string{"user"}, // Role padrão
		Scopes:        []string{},
		VerifiedEmail: false,
	}
}

// HasRole verifica se o usuário possui um determinado role
func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// AddRole adiciona um novo role ao usuário
func (u *User) AddRole(role string) {
	if !u.HasRole(role) {
		u.Roles = append(u.Roles, role)
	}
}

// RemoveRole remove um role do usuário
func (u *User) RemoveRole(role string) {
	for i, r := range u.Roles {
		if r == role {
			u.Roles = append(u.Roles[:i], u.Roles[i+1:]...)
			break
		}
	}
}

// HasScope verifica se o usuário possui um determinado escopo
func (u *User) HasScope(scope string) bool {
	for _, s := range u.Scopes {
		if s == scope {
			return true
		}
	}
	return false
}

// AddScope adiciona um novo escopo ao usuário
func (u *User) AddScope(scope string) {
	if !u.HasScope(scope) {
		u.Scopes = append(u.Scopes, scope)
	}
}

// RemoveScope remove um escopo do usuário
func (u *User) RemoveScope(scope string) {
	for i, s := range u.Scopes {
		if s == scope {
			u.Scopes = append(u.Scopes[:i], u.Scopes[i+1:]...)
			break
		}
	}
}

// UpdateOAuthInfo atualiza as informações do OAuth do usuário
func (u *User) UpdateOAuthInfo(accessToken, refreshToken string, tokenExpiry int64, scopes []string) {
	u.AccessToken = accessToken
	u.RefreshToken = refreshToken
	u.TokenExpiry = tokenExpiry
	u.Scopes = scopes
}

// UpdateProfile atualiza as informações do perfil do usuário
func (u *User) UpdateProfile(name, givenName, familyName, picture string, verifiedEmail bool) {
	u.Name = name
	u.GivenName = givenName
	u.FamilyName = familyName
	u.Picture = picture
	u.VerifiedEmail = verifiedEmail
}
