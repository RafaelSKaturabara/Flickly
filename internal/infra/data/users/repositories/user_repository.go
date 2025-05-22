package repositories

import (
	"context"
	"errors"
	"flickly/internal/domain/users/entities"

	"github.com/google/uuid"
)

type UserRepository struct {
	Users []entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	for _, existingUser := range r.Users {
		if user.Email == existingUser.Email {
			return errors.New("user already exists")
		}
	}
	r.Users = append(r.Users, *user)
	return nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	for _, user := range r.Users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	for _, user := range r.Users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	for i, u := range r.Users {
		if u.ID == user.ID {
			r.Users[i] = *user
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	for i, user := range r.Users {
		if user.ID == id {
			r.Users = append(r.Users[:i], r.Users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) GetUserByProviderID(ctx context.Context, provider, providerID string) (*entities.User, error) {
	for _, user := range r.Users {
		if user.Provider == provider && user.ProviderID == providerID {
			return &user, nil
		}
	}
	return nil, nil
}

func (r *UserRepository) UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error {
	for i, user := range r.Users {
		if user.ID == userID {
			r.Users[i].AccessToken = accessToken
			r.Users[i].RefreshToken = refreshToken
			r.Users[i].TokenExpiry = tokenExpiry
			r.Users[i].Scopes = scopes
			return nil
		}
	}
	return errors.New("user not found")
}

func (r *UserRepository) UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error {
	for i, user := range r.Users {
		if user.ID == userID {
			r.Users[i].Roles = roles
			return nil
		}
	}
	return errors.New("user not found")
}
