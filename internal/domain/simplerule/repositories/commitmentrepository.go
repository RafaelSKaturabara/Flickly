package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
)

type ICommitmentRepository interface {
	core.Repository[entities.Commitment]
}
