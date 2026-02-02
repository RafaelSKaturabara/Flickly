package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type RoleRespository struct {
	core.GormRepository[entities.Role, gormmappings.RoleDB]
}

func NewRoleRespository(db *gorm.DB) domainRepositories.IRoleRepository {
	return &RoleRespository{
		core.GormRepository[entities.Role, gormmappings.RoleDB](*core.NewGormRepository[entities.Role, gormmappings.RoleDB](db)),
	}
}
	