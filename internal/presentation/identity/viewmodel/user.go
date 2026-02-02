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
	Expires_in int    `json:"expires_in"`
}
