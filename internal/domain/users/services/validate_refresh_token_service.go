package services

import (
	"context"
	"fmt"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/google/uuid"
)

type ValidateRefreshTokenService struct {
}

func NewValidateRefreshTokenService() *ValidateRefreshTokenService {
	return &ValidateRefreshTokenService{}
}

func (s *ValidateRefreshTokenService) AbleToRun(ctx context.Context, entity core.Entity) bool {
	return true
}

func (s *ValidateRefreshTokenService) Run(ctx context.Context, entity core.Entity) error {
	return nil
}

func (s *ValidateRefreshTokenService) ValidateRefreshToken(tokenString string) (uuid.UUID, error) {
	claims, err := utilities.ValidateToken(tokenString, "s.jwtSecret")
	if err != nil {
		return uuid.Nil, err
	}

	if userIDStr, ok := claims["user_id"].(string); ok {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return uuid.Nil, err
		}
		return userID, nil
	}

	return uuid.Nil, fmt.Errorf("token inválido: user_id não encontrado")
}
