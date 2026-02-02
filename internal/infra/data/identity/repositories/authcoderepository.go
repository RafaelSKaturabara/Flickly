package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type AuthCodeRepository struct {
	core.GormRepository[entities.AuthCode, gormmappings.AuthCodeDB]
}

func NewAuthCodeRepository(db *gorm.DB) domainRepositories.IAuthCodeRepository {
	return &AuthCodeRepository{
		core.GormRepository[entities.AuthCode, gormmappings.AuthCodeDB](*core.NewGormRepository[entities.AuthCode, gormmappings.AuthCodeDB](db)),
	}
}
