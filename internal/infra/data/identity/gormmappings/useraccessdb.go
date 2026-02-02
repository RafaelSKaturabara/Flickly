package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserAccessDB struct {
	core.EntityDB `gorm:"embedded"`
	UserID        uuid.UUID `gorm:"type:uuid;column:user_id;not null"`
	User          *UserDB   `gorm:"foreignKey:UserID"`
	ClientID      uuid.UUID `gorm:"type:uuid;column:client_id;not null"`
	Client        *ClientDB `gorm:"foreignKey:ClientID"`
	RolesDB       []RoleDB  `gorm:"many2many:user_access_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
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
