package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
)

type GenerateRefreshTokenService struct {
}

func NewGenerateRefreshTokenService() *GenerateRefreshTokenService {
	return &GenerateRefreshTokenService{}
}

func (s *GenerateRefreshTokenService) AbleToRun(ctx context.Context, entity core.Entity) bool {
	return entity.GetID() != uuid.Nil
}

func (s *GenerateRefreshTokenService) Run(ctx context.Context, entity core.Entity) error {
	user := entity.(*entities.User)
	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return err
	}

	user.RefreshToken = refreshToken
	return nil
}

func (s *GenerateRefreshTokenService) generateRefreshToken(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.GetID().String(),
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Token válido por 7 dias
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("s.jwtSecret"))
}
