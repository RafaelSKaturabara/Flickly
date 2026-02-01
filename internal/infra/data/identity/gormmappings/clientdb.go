package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ClientDB struct {
	core.BaseEntity `gorm:"embedded"`
	Name            string          `gorm:"column:name;size:255;not null"`
	Secret          string          `gorm:"column:secret;size:255;not null"`
	RedirectURIs    []RedirectURIDB `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (ClientDB) TableName() string {
	return "clients"
}

func (db *ClientDB) ToDomain() entities.Client {
	var d entities.Client
	_ = copier.Copy(&d, db)
	return d
}

func (db *ClientDB) FromDomain(d entities.Client) {
	_ = copier.Copy(db, &d)
}

func SeedClients(db *gorm.DB) {}
