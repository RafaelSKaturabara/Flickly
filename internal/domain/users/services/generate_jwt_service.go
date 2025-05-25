package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
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
	token, err := s.generateJWT(user)
	if err != nil {
		return err
	}

	user.AccessToken = token
	user.TokenType = "Bearer"
	return nil
}

func (s *GenerateJWTService) generateJWT(user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.GetID().String(),
		"email": user.Email,
		"name":  user.Name,
		"roles": user.Roles,
		"exp":   time.Now().Add(time.Minute * 15).Unix(), // Token válido por 15 minutos
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("config.JWTSecret"))
}
