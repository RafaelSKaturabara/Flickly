package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/valueobjects"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserDB struct {
	core.EntityDB `gorm:"embedded"`
	Email         string                    `gorm:"column:email;uniqueIndex;size:255;not null"`
	Name          string                    `gorm:"column:name;size:1024;not null"`
	PasswordHash  valueobjects.PasswordHash `gorm:"column:password_hash"`
	Picture       string                    `gorm:"column:picture;size:1024"`
	VerifiedEmail bool                      `gorm:"column:verified_email;default:false"`
	AccessToken   string                    `gorm:"column:access_token;size:2048"`
	UserAccesses  []UserAccessDB            `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
