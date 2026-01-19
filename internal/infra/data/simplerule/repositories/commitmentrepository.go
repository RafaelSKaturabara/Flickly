package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
	"github.com/rkaturabara/flickly/internal/infra/data/core"
	"gorm.io/gorm"
)

type CommitmentRepository struct {
	core.GormRepository[entities.Commitment]
}

func NewCommitmentRepository(dbe *gorm.DB) *CommitmentRepository {
	return &CommitmentRepository{
		core.GormRepository[entities.Commitment](*core.NewGormRepository[entities.Commitment](dbe)),
	}
} 