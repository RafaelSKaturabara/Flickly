package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RedirectURIDB struct {
	entities.RedirectURI `gorm:"embedded"`
	Value           string    `gorm:"column:value;size:2048;not null"`
	ClientID        uuid.UUID `gorm:"column:client_id;type:uuid;not null"`
	Client          *ClientDB `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RedirectURIDB) TableName() string {
	return "redirect_uris"
}

func (db *RedirectURIDB) ToDomain() entities.RedirectURI {
	var d entities.RedirectURI
	_ = copier.Copy(&d, db)
	return d
}

func (db *RedirectURIDB) FromDomain(d entities.RedirectURI) {
	_ = copier.Copy(db, &d)
}

func SeedRedirectURIs(db *gorm.DB) {}
