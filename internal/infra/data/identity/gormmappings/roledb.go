package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RoleDB struct {
	core.EntityDB `gorm:"embedded"`
	Value         string     `gorm:"column:value;size:255;not null"`
	ClientID      *uuid.UUID `gorm:"type:uuid;column:client_id"`
	Client        *ClientDB  `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (RoleDB) TableName() string {
	return "roles"
}

func (db *RoleDB) ToDomain() entities.Role {
	var d entities.Role
	_ = copier.Copy(&d, db)
	return d
}

func (db *RoleDB) FromDomain(d entities.Role) {
	_ = copier.Copy(db, &d)
}

func SeedRoles(db *gorm.DB) {
	var count int64
	db.Model(&RoleDB{}).Count(&count)
	if count > 0 {
		return // Já tem dados
	}

	var admin, user entities.Role
	var adminDB, userDB RoleDB

	admin = entities.NewRole(entities.RoleAdmin)
	user = entities.NewRole(entities.RoleUser)

	adminDB.FromDomain(admin)
	userDB.FromDomain(user)

	roles := []RoleDB{
		adminDB,
		userDB,
	}

	for _, role := range roles {
		db.Create(&role)
	}
}
