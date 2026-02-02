package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type AccessGrantDB struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid"`
	entities.AccessGrant `gorm:"embedded"`
	Client               *ClientDB `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User                 *UserDB   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
