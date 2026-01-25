package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type UserDB struct {
	core.BaseEntityDB `gorm:"embedded"`
	Email             string   `gorm:"column:email;size:255;unique;not null"`
	Name              string   `gorm:"column:name;size:255;not null"`
	Password          string   `gorm:"column:password;size:255;not null"`
	Picture           string   `gorm:"column:picture;size:1024"`
	VerifiedEmail     bool     `gorm:"column:verified_email"`
	Roles             []string `gorm:"column:roles;serializer:json"`
	AccessToken       string   `gorm:"column:access_token;size:2048"`
	TokenType         string   `gorm:"column:token_type;size:50"`
	TokenExpiry       int64    `gorm:"column:token_expiry"`
	TokenScopes       []string `gorm:"column:token_scopes;serializer:json"`
	ClientID          string   `gorm:"column:client_id;size:255"`
	ClientSecret      string   `gorm:"column:client_secret;size:255"`
}

func (UserDB) TableName() string {
	return "users"
}

func (db *UserDB) ToDomain() entities.User {
	return entities.User{
		BaseEntity:    db.BaseEntityDB.BaseEntity,
		Email:         db.Email,
		Name:          db.Name,
		Password:      db.Password,
		Picture:       db.Picture,
		VerifiedEmail: db.VerifiedEmail,
		Roles:         db.Roles,
		AccessToken:   db.AccessToken,
		TokenType:     db.TokenType,
		TokenExpiry:   db.TokenExpiry,
		TokenScopes:   db.TokenScopes,
		ClientID:      db.ClientID,
		ClientSecret:  db.ClientSecret,
	}
}

func (db *UserDB) FromDomain(d entities.User) {
	db.BaseEntityDB.BaseEntity = d.BaseEntity
	db.Email = d.Email
	db.Name = d.Name
	db.Password = d.Password
	db.Picture = d.Picture
	db.VerifiedEmail = d.VerifiedEmail
	db.Roles = d.Roles
	db.AccessToken = d.AccessToken
	db.TokenType = d.TokenType
	db.TokenExpiry = d.TokenExpiry
	db.TokenScopes = d.TokenScopes
	db.ClientID = d.ClientID
	db.ClientSecret = d.ClientSecret
}

func SeedUsers(db *gorm.DB) {}
