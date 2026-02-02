package gormmappings

import (
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AuthCodeDB struct {
	core.EntityDB `gorm:"embedded"`
	Code          string     `gorm:"column:code;size:255;not null"`
	ClientID      uuid.UUID  `gorm:"type:uuid;column:client_id;not null"`
	Client        *ClientDB  `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserID        *uuid.UUID `gorm:"type:uuid;column:user_id"`
	User          *UserDB    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CodeChallenge string     `gorm:"column:code_challenge;size:255"`
	ExpiresAt     time.Time  `gorm:"column:expires_at;not null"`
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
