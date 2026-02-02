package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"
)

type AuthCodeDB struct {
	entities.AuthCode `gorm:"embedded"`
	Code          string    `gorm:"column:code;size:255;not null;uniqueIndex"`
	ClientID      uuid.UUID `gorm:"column:client_id;type:uuid;not null"`
	Client        *ClientDB `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	UserID        uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	User          *UserDB   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	CodeChallenge string    `gorm:"column:code_challenge;size:255;not null"`
	ExpiresAt     time.Time `gorm:"column:expires_at;not null;index"`
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
