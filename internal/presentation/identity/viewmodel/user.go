package viewmodel

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetAuthorizeResponse struct {
	Code       string `json:"code"`
	ClientName string `json:"client_name"`
	ExpiresIn  int    `json:"expires_in"`
}

type CreateLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Code     string `json:"code" binding:"required,uuid"`
}

type CreateLoginResponse struct {
	Code string `json:"code" binding:"required,uuid"`
}
