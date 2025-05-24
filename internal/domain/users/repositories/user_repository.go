package repositories

import (
	"context"

	"github.com/rkaturabara/flickly/internal/domain/users/entities"

	"github.com/google/uuid"
)

type IUserRepository interface {
	// Métodos básicos
	CreateUser(ctx context.Context, user *entities.User) error
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	UpdateUser(ctx context.Context, user *entities.User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error

	// Métodos específicos para OAuth2
	GetUserByProviderID(ctx context.Context, provider, providerID string) (*entities.User, error)
	UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error
	UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error
}
