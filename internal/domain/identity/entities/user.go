package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/valueobjects"
)

// User representa um usuário no sistema com suporte a OAuth2
type User struct {
	core.BaseEntity
	Email         string
	Name          string
	PasswordHash  valueobjects.PasswordHash
	Picture       string
	VerifiedEmail bool
	UserAccesses  []UserAccess
	AccessToken   string
}

// NewUser cria uma nova instância de User
func NewUser(name, email string, passwordHash valueobjects.PasswordHash, userAccesses []UserAccess) *User {
	return &User{
		BaseEntity:    core.NewBaseEntity(),
		Email:         email,
		Name:          name,
		UserAccesses:  userAccesses,
		VerifiedEmail: false,
		PasswordHash:  passwordHash,
	}
}

// UpdateProfile atualiza as informações do perfil do usuário
func (u *User) UpdateProfile(name, givenName, familyName, picture string, verifiedEmail bool) {
	u.Name = name
	u.Picture = picture
	u.VerifiedEmail = verifiedEmail
}

func (u *User) IsValid() bool {
	return true
}
