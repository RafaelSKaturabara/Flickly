package entities

import (
	"github.com/rkaturabara/flickly/internal/domain/core"
)

// User representa um usuário no sistema com suporte a OAuth2
type User struct {
	core.BaseEntity
	Email         string   `json:"email"`
	Name          string   `json:"name"`
	Password      string   `json:"-"` // Senha não é serializada em JSON
	Picture       string   `json:"picture,omitempty"`
	VerifiedEmail bool     `json:"verified_email"`
	Roles         []string `json:"roles"`
	AccessToken   string   `json:"-"`
	RefreshToken  string   `json:"-"`
	TokenExpiry   int64    `json:"-"`
	TokenScopes   []string `json:"-"`
	ClientID      string   `json:"client_id"`
	ClientSecret  string   `json:"client_secret"`
}

// NewUser cria uma nova instância de User
func NewUser(name, email, clientID, clientSecret, password string) *User {
	return &User{
		BaseEntity:    core.NewBaseEntity(),
		Email:         email,
		Name:          name,
		Roles:         []string{"user"},
		VerifiedEmail: false,
		ClientID:      clientID,
		ClientSecret:  clientSecret,
		Password:      password,
	}
}

// HasRole verifica se o usuário possui uma determinada role
func (u *User) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// AddRole adiciona uma role ao usuário
func (u *User) AddRole(role string) {
	if !u.HasRole(role) {
		u.Roles = append(u.Roles, role)
	}
}

// RemoveRole remove uma role do usuário
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
	for _, s := range u.TokenScopes {
		if s == scope {
			return true
		}
	}
	return false
}

// AddScope adiciona um escopo ao usuário
func (u *User) AddScope(scope string) {
	if !u.HasScope(scope) {
		u.TokenScopes = append(u.TokenScopes, scope)
	}
}

// RemoveScope remove um escopo do usuário
func (u *User) RemoveScope(scope string) {
	for i, s := range u.TokenScopes {
		if s == scope {
			u.TokenScopes = append(u.TokenScopes[:i], u.TokenScopes[i+1:]...)
			break
		}
	}
}

// UpdateOAuthInfo atualiza as informações de OAuth do usuário
func (u *User) UpdateOAuthInfo(accessToken, refreshToken string, tokenExpiry int64, scopes []string) {
	u.AccessToken = accessToken
	u.RefreshToken = refreshToken
	u.TokenExpiry = tokenExpiry
	u.TokenScopes = scopes
}

// UpdateProfile atualiza as informações do perfil do usuário
func (u *User) UpdateProfile(name, givenName, familyName, picture string, verifiedEmail bool) {
	u.Name = name
	u.Picture = picture
	u.VerifiedEmail = verifiedEmail
}
