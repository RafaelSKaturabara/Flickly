package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
)

type IAccessGrantRepository interface {
	core.Repository[entities.AccessGrant]
	GetByTokenHash(tokenHash string) (*entities.AccessGrant, error)
}