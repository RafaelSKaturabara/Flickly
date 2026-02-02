package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserDB struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid"`
	entities.User `gorm:"embedded"`
	Email         string         `gorm:"column:email;uniqueIndex;size:255;not null"`
	Name          string         `gorm:"column:name;size:1024;not null"`
	UserAccesses  []UserAccessDB `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
