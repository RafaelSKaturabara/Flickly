package repositories

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
)

type IRedirectURIRepository interface {
	core.Repository[entities.RedirectURI]
}
