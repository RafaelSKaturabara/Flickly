package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type RedirectURIRepository struct {
	core.GormRepository[entities.RedirectURI, gormmappings.RedirectURIDB]
}

func NewRedirectURIRepository(db *gorm.DB) domainRepositories.IRedirectURIRepository {
	return &RedirectURIRepository{
		core.GormRepository[entities.RedirectURI, gormmappings.RedirectURIDB](*core.NewGormRepository[entities.RedirectURI, gormmappings.RedirectURIDB](db)),
	}
}