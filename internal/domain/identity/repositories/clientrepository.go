package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
)

type IClientRepository interface {
	core.Repository[entities.Client]
}
