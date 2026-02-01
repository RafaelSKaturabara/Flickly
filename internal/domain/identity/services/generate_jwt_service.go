package services

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/google/uuid"
)

type GenerateJWTService struct {
}

func NewGenerateJWTService() *GenerateJWTService {
	return &GenerateJWTService{}
}

func (s *GenerateJWTService) AbleToRun(ctx context.Context, entity core.Entity) bool {
	return entity.GetID() != uuid.Nil
}

func (s *GenerateJWTService) Run(ctx context.Context, entity core.Entity) error {
	user := entity.(*entities.User)

	claims := map[string]interface{}{
		"id":    user.GetID().String(),
		"email": user.Email,
		"name":  user.Name,
		"roles": user.Roles,
	}

	token, err := utilities.GenerateToken("config.JWTSecret", 15, claims)
	if err != nil {
		return err
	}

	user.AccessToken = token
	user.TokenType = "Bearer"
	return nil
}
