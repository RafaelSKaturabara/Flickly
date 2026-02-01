package repositories

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"

	"github.com/google/uuid"
)

type IUserRepository interface {
	core.Repository[entities.User]
	// Métodos básicos
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByEmailAndPasswordAndClientAndSecret(ctx context.Context, email, password string, clientID string, clientSecret string) (*entities.User, error)

	// Métodos específicos para OAuth2
	UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error
	UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error
}
