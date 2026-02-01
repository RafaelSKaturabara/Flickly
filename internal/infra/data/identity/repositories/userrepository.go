package repositories

import (
	"context"
	domainRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"gorm.io/gorm"
)

type UserRepository struct {
	core.GormRepository[entities.User, gormmappings.UserDB]
}

func NewUserRepository(db *gorm.DB) domainRepositories.IUserRepository {
	return &UserRepository{
		core.GormRepository[entities.User, gormmappings.UserDB](*core.NewGormRepository[entities.User, gormmappings.UserDB](db)),
	}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var result entities.User
	if err := r.DB.Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepository) GetUserByEmailAndPasswordAndClient(ctx context.Context, email, password, clientID string) (*entities.User, error) {
	var result entities.User

	panic("query com client id")
	if err := r.DB.Where("email = ? AND password = ? AND client_id = ?", email, password, clientID).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
