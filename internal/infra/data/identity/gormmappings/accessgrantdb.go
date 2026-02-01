package gormmappings

import (
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AccessGrantDB struct {
	core.BaseEntity `gorm:"embedded"`
	UserID          uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	ClientID        uuid.UUID `gorm:"column:client_id;type:uuid;not null"`
	TokenHash       string    `gorm:"column:token_hash;size:255;not null"`
	ExpiresAt       time.Time `gorm:"column:expires_at;not null"`
	Revoked         bool      `gorm:"column:revoked;default:false"`
	UserAgent       string    `gorm:"column:user_agent;size:1024"`
	IPAddress       string    `gorm:"column:ip_address;size:255"`
	Client          *ClientDB `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User            *UserDB   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
