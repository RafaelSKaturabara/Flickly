package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rkaturabara/flickly/internal/application/users/viewmodel"
	"github.com/rkaturabara/flickly/internal/domain/core"
	"github.com/rkaturabara/flickly/internal/domain/users/entities"
)

type contextKey string

const UserContextKey contextKey = "user"

type JWTMiddleware struct {
	jwtSecret string
}

func NewJWTMiddleware(jwtSecret string) *JWTMiddleware {
	return &JWTMiddleware{
		jwtSecret: jwtSecret,
	}
}

// Auth é o middleware base de autenticação
func (m *JWTMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			name := claims["name"].(string)
			id := claims["id"].(string)
			roles := make([]string, 0)
			if rolesClaim, ok := claims["roles"].([]interface{}); ok {
				for _, role := range rolesClaim {
					roles = append(roles, role.(string))
				}
			}

			user := &entities.User{
				Email: email,
				Name:  name,
				Roles: roles,
				BaseEntity: core.BaseEntity{
					ID: uuid.MustParse(id),
				},
			}

			// Adiciona o usuário ao context.Context
			ctx := context.WithValue(c.Request.Context(), UserContextKey, user)
			c.Request = c.Request.WithContext(ctx)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
		}
	}
}

// Auth é o middleware base de autenticação
func (m *JWTMiddleware) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenRequest viewmodel.TokenRequest
		if err := c.ShouldBind(&tokenRequest); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		if tokenRequest.GrantType != "refresh_token" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(m.jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			name := claims["name"].(string)
			id := claims["id"].(string)
			roles := make([]string, 0)
			if rolesClaim, ok := claims["roles"].([]interface{}); ok {
				for _, role := range rolesClaim {
					roles = append(roles, role.(string))
				}
			}

			user := &entities.User{
				Email: email,
				Name:  name,
				Roles: roles,
				BaseEntity: core.BaseEntity{
					ID: uuid.MustParse(id),
				},
			}

			// Adiciona o usuário ao context.Context
			ctx := context.WithValue(c.Request.Context(), UserContextKey, user)
			c.Request = c.Request.WithContext(ctx)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
		}
	}
}

// Role é um middleware que verifica se o usuário tem a role necessária
func (m *JWTMiddleware) Role(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Request.Context().Value(UserContextKey).(*entities.User)
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
			c.Abort()
			return
		}

		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar usuário"})
			c.Abort()
			return
		}

		for _, userRole := range user.Roles {
			if userRole == role {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado"})
		c.Abort()
	}
}
