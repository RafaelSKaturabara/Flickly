package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type ClientDB struct {
	core.EntityDB `gorm:"embedded"`
	Secret        string          `gorm:"column:secret;size:255;not null"`
	Name          string          `gorm:"column:name;size:255;not null"`
	RedirectURIs  []RedirectURIDB `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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

func SeedClients(db *gorm.DB) {
	var count int64
	db.Model(&ClientDB{}).Count(&count)
	if count > 0 {
		return // Já tem dados
	}

	clientDB := ClientDB{}

	client := entities.NewClient("Simle Rule", "26952b22-5a37-4c84-a96c-4ba6ef0c14e6")
	client.ID = uuid.MustParse("8e705001-1089-4e90-86eb-5c10aa609165")

	clientDB.FromDomain(client)

	db.Create(&clientDB)
}
