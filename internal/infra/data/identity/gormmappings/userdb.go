package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/valueobjects"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserDB struct {
	core.BaseEntity `gorm:"embedded"`
	Email           string                    `gorm:"column:email;uniqueIndex;size:255;not null"`
	Name            string                    `gorm:"column:name;size:1024;not null"`
	PasswordHash    valueobjects.PasswordHash `gorm:"column:password_hash"`
	Picture         string                    `gorm:"column:picture;size:2048"`
	VerifiedEmail   bool                      `gorm:"column:verified_email;default:false"`
	UserAccesses    []UserAccessDB            `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccessToken     string                    `gorm:"-"`
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
