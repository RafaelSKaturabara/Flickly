package gormmappings

import (
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AccessGrantDB struct {
	core.EntityDB `gorm:"embedded"`
	UserID        uuid.UUID `gorm:"type:uuid;column:user_id;not null"`
	User          *UserDB   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ClientID      uuid.UUID `gorm:"type:uuid;column:client_id;not null"`
	Client        *ClientDB `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TokenHash     string    `gorm:"column:token_hash;size:255;not null"`
	ExpiresAt     time.Time `gorm:"column:expires_at;not null"`
	Revoked       bool      `gorm:"column:revoked;default:false"`
	UserAgent     string    `gorm:"column:user_agent;size:1024"`
	IPAddress     string    `gorm:"column:ip_address;size:255"`
}

func (AccessGrantDB) TableName() string {
	return "access_grants"
}

func (db *AccessGrantDB) ToDomain() entities.AccessGrant {
	var d entities.AccessGrant
	_ = copier.Copy(&d, db)
	return d
}

func (db *AccessGrantDB) FromDomain(d entities.AccessGrant) {
	_ = copier.Copy(db, &d)
}

func SeedAccessGrants(db *gorm.DB) {}
