package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserAccessDB struct {
	ID                  uuid.UUID `gorm:"primaryKey;type:uuid"`
	entities.UserAccess `gorm:"embedded"`
	User                *UserDB   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Client              *ClientDB `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Roles               []RoleDB  `gorm:"many2many:user_access_roles;foreignKey:ID;joinForeignKey:UserAccessID;references:ID;joinReferences:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (UserAccessDB) TableName() string {
	return "user_accesses"
}

func (db *UserAccessDB) ToDomain() entities.UserAccess {
	var d entities.UserAccess
	_ = copier.Copy(&d, db)
	return d
}

func (db *UserAccessDB) FromDomain(d entities.UserAccess) {
	_ = copier.Copy(db, &d)
}

func SeedUserAccess(db *gorm.DB) {}
