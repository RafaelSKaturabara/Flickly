package services

var dfsdf = 0

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/gofrs/uuid"
// 	"github.com/rkaturabara/flickly/internal/domain/core"
// 	"github.com/rkaturabara/flickly/internal/domain/users/entities"
// 	"github.com/rkaturabara/flickly/internal/domain/users/repositories"

// 	"github.com/golang-jwt/jwt/v5"
// 	"golang.org/x/crypto/bcrypt"
// )

// type AuthServiceInterface interface {
// 	Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error)
// 	Token(ctx context.Context, req TokenRequest) (*TokenResponse, error)
// }

// type AuthService struct {
// 	userRepository repositories.IUserRepository
// 	jwtSecret      string
// 	clientID       string
// 	clientSecret   string
// }

// func NewAuthService(userRepository repositories.IUserRepository, jwtSecret, clientID, clientSecret string) *AuthService {
// 	return &AuthService{
// 		userRepository: userRepository,
// 		jwtSecret:      jwtSecret,
// 		clientID:       clientID,
// 		clientSecret:   clientSecret,
// 	}
// }

// func (s *AuthService) Register(ctx context.Context, req RegisterRequest) (*RegisterResponse, error) {
// 	// Verifica se o usuário já existe
// 	existingUser, err := s.userRepository.GetUserByEmail(ctx, req.Email)
// 	if err == nil && existingUser != nil {
// 		return nil, repositories.ErrUserAlreadyExists
// 	}

// 	// Hash da senha
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Cria o usuário
// 	user := entities.NewUser(req.Name, req.Email, "local", req.Email)
// 	user.Password = string(hashedPassword)

// 	// Salva o usuário
// 	if err := s.userRepository.CreateUser(ctx, user); err != nil {
// 		return nil, err
// 	}

// 	return &RegisterResponse{
// 		User: user,
// 	}, nil
// }

// func (s *AuthService) Token(ctx context.Context, req TokenRequest) (*TokenResponse, error) {
// 	// Verifica o client_id e client_secret
// 	if req.ClientID != s.clientID || req.ClientSecret != s.clientSecret {
// 		return nil, ErrInvalidClient
// 	}

// 	// Verifica o grant_type
// 	if req.GrantType != "password" {
// 		return nil, ErrInvalidGrant
// 	}

// 	// Busca o usuário pelo email
// 	user, err := s.userRepository.GetUserByEmail(ctx, req.Username)
// 	if err != nil {
// 		return nil, ErrInvalidCredentials
// 	}

// 	// Verifica a senha
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
// 		return nil, ErrInvalidCredentials
// 	}

// 	// Gera o token JWT
// 	token, err := s.generateJWT(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Gera o refresh token
// 	refreshToken, err := s.generateRefreshToken(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &TokenResponse{
// 		AccessToken:  token,
// 		TokenType:    "Bearer",
// 		ExpiresIn:    3600, // 1 hora
// 		RefreshToken: refreshToken,
// 		Scope:        req.Scope,
// 	}, nil
// }

// func (s *AuthService) generateRefreshToken(user *entities.User) (string, error) {
// 	claims := jwt.MapClaims{
// 		"user_id": user.GetID().String(),
// 		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Token válido por 7 dias
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(s.jwtSecret))
// }
