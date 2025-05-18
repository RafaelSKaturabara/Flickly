package view_models

import (
	"github.com/google/uuid"
	"time"
)

type CreatePessoaResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Nome      string    `json:"nome"`
	Idade     int       `json:"idade"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
