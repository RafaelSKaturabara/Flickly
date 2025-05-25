package services

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/domain/core"
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
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte("s.jwtSecret"), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return uuid.Nil, err
		}
		return userID, nil
	}

	return uuid.Nil, fmt.Errorf("token inválido")
}
