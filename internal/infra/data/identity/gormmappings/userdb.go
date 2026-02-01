package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserDB struct {
	entities.User `gorm:"embedded"`
	Email         string          `gorm:"column:email;size:255;unique;not null"`
	Name          string          `gorm:"column:name;size:255;not null"`
	Password      string          `gorm:"column:password;size:255;not null"`
	Picture       string          `gorm:"column:picture;size:1024"`
	VerifiedEmail bool            `gorm:"column:verified_email"`
	Roles         []entities.Role `gorm:"column:roles;type:text;serializer:json"`
	AccessToken   string          `gorm:"column:access_token;size:2048"`
	TokenType     string          `gorm:"column:token_type;size:50"`
	TokenExpiry   int64           `gorm:"column:token_expiry"`
	TokenScopes   []string        `gorm:"column:token_scopes;type:text;serializer:json"`
	ClientID      string          `gorm:"column:client_id;size:255"`
	ClientSecret  string          `gorm:"column:client_secret;size:255"`
}

func (UserDB) TableName() string {
	return "users"
}

func (db *UserDB) ToDomain() entities.User {
	var d entities.User
	_ = copier.Copy(&d, db)
	return d
}

func (db *UserDB) FromDomain(d entities.User) {
	_ = copier.Copy(db, &d)
}

func SeedUsers(db *gorm.DB) {}
