package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	"gorm.io/gorm"
)

type CommitmentRepository struct {
	core.GormRepository[entities.Commitment, gormmappings.CommitmentDB]
}

func NewCommitmentRepository(dbe *gorm.DB) domainRepositories.ICommitmentRepository {
	return &CommitmentRepository{
		core.GormRepository[entities.Commitment, gormmappings.CommitmentDB](*core.NewGormRepository[entities.Commitment, gormmappings.CommitmentDB](dbe)),
	}
} 
