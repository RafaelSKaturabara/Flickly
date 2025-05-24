package user_repository

import (
	"context"

	"sync"

	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
)

type UserRepository struct {
	users map[uuid.UUID]*entities.User
	mu    sync.RWMutex
}

func NewUserRepository() repositories.IUserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]*entities.User),
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Verifica se já existe um usuário com o mesmo email
	for _, existingUser := range r.users {
		if existingUser.Email == user.Email {
			return repositories.ErrUserAlreadyExists
		}
	}

	r.users[user.GetID()] = user
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, repositories.ErrUserNotFound
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, repositories.ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.GetID()]; !exists {
		return repositories.ErrUserNotFound
	}

	r.users[user.GetID()] = user
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return repositories.ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}

func (r *UserRepository) GetUserByProviderID(ctx context.Context, provider, providerID string) (*entities.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Provider == provider && user.ProviderID == providerID {
			return user, nil
		}
	}

	return nil, repositories.ErrUserNotFound
}

func (r *UserRepository) UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[userID]
	if !exists {
		return repositories.ErrUserNotFound
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	user.TokenExpiry = tokenExpiry
	user.TokenScopes = scopes

	r.users[userID] = user
	return nil
}

func (r *UserRepository) UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[userID]
	if !exists {
		return repositories.ErrUserNotFound
	}

	user.Roles = roles
	r.users[userID] = user
	return nil
}
