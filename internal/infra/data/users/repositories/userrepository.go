package repositories

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/users/gormmappings"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	core.GormRepository[entities.User, gormmappings.UserDB]
}

func NewUserRepository(dbe *gorm.DB) *UserRepository {
	return &UserRepository{
		core.GormRepository[entities.User, gormmappings.UserDB](*core.NewGormRepository[entities.User, gormmappings.UserDB](dbe)),
	}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var model gormmappings.UserDB
	err := r.DB.WithContext(ctx).First(&model, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repositories.ErrUserNotFound
		}
		return nil, err
	}
	domain := model.ToDomain()
	return &domain, nil
}

func (r *UserRepository) GetUserByEmailAndPasswordAndClientAndSecret(ctx context.Context, email, password, clientID, clientSecret string) (*entities.User, error) {
	var model gormmappings.UserDB
	err := r.DB.WithContext(ctx).First(&model, "email = ? AND client_id = ? AND client_secret = ?", email, clientID, clientSecret).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, repositories.ErrUserNotFound
		}
		return nil, err
	}

	if utilities.CompareHashAndPassword(model.Password, password) != nil {
		return nil, repositories.ErrUserNotFound
	}

	domain := model.ToDomain()
	return &domain, nil
}

func (r *UserRepository) UpdateUserOAuthInfo(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, tokenExpiry int64, scopes []string) error {
	return r.DB.WithContext(ctx).Model(&gormmappings.UserDB{}).Where("id = ?", userID).Updates(&gormmappings.UserDB{
		AccessToken: accessToken,
		TokenExpiry: tokenExpiry,
		TokenScopes: scopes,
	}).Error
}

func (r *UserRepository) UpdateUserRoles(ctx context.Context, userID uuid.UUID, roles []string) error {
	return r.DB.WithContext(ctx).Model(&gormmappings.UserDB{}).Where("id = ?", userID).Updates(&gormmappings.UserDB{
		Roles: roles,
	}).Error
}
