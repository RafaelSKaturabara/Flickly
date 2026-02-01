package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type AccessGrantRepository struct {
	core.GormRepository[entities.AccessGrant, gormmappings.AccessGrantDB]
}

func NewAccessGrantRepository(db *gorm.DB) domainRepositories.IAccessGrantRepository {
	return &AccessGrantRepository{
		core.GormRepository[entities.AccessGrant, gormmappings.AccessGrantDB](*core.NewGormRepository[entities.AccessGrant, gormmappings.AccessGrantDB](db)),
	}
}

func (r *AccessGrantRepository) GetByTokenHash(tokenHash string) (*entities.AccessGrant, error) {
	var result entities.AccessGrant
	if err := r.DB.Where("token_hash = ?", tokenHash).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
