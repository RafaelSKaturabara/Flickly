package repositories

import (
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/simplerule/entities"
)

type ICommitmentRepository interface {
	core.Repository[entities.Commitment]
}
