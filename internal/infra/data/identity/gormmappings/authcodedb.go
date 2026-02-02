package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AuthCodeDB struct {
	ID                uuid.UUID `gorm:"primaryKey;type:uuid"`
	entities.AuthCode `gorm:"embedded"`
	Client            *ClientDB `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User              *UserDB   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (AuthCodeDB) TableName() string {
	return "auth_codes"
}

func (db *AuthCodeDB) ToDomain() entities.AuthCode {
	var d entities.AuthCode
	_ = copier.Copy(&d, db)
	return d
}

func (db *AuthCodeDB) FromDomain(d entities.AuthCode) {
	_ = copier.Copy(db, &d)
}

func SeedAuthCodes(db *gorm.DB) {}
