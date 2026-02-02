package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type UserAccessRepository struct {
	core.GormRepository[entities.UserAccess, gormmappings.UserAccessDB]
}

func NewUserAccessRepository(db *gorm.DB) domainRepositories.IUserAccessRepository {
	return &UserAccessRepository{
		core.GormRepository[entities.UserAccess, gormmappings.UserAccessDB](*core.NewGormRepository[entities.UserAccess, gormmappings.UserAccessDB](db)),
	}
}
