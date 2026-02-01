package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserAccessDB struct {
	core.BaseEntity `gorm:"embedded"`
	UserID          uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	ClientID        uuid.UUID `gorm:"column:client_id;type:uuid;not null"`
	User            *UserDB   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Client          *ClientDB `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Roles           []RoleDB  `gorm:"many2many:user_access_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
