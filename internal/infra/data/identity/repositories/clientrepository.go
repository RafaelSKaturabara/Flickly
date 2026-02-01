package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type ClientRepository struct {
	core.GormRepository[entities.Client, gormmappings.ClientDB]
}

func NewClientRepository(db *gorm.DB) domainRepositories.IClientRepository {
	return &ClientRepository{
		core.GormRepository[entities.Client, gormmappings.ClientDB](*core.NewGormRepository[entities.Client, gormmappings.ClientDB](db)),
	}
}
