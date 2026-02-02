package repositories

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
)

type IUserRepository interface {
	core.Repository[entities.User]
	// Métodos básicos
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByEmailAndPasswordAndClient(ctx context.Context, email, passwordHash, clientID string) (*entities.User, error)
}
