package view_models

import (
	"github.com/google/uuid"
	"time"
)

type GetPessoaResponse struct {
	ID           uuid.UUID  `json:"id"`
	CreatedAt    time.Time  `json:"createdAt"`
	LastUpdateAt *time.Time `json:"lastUpdateAt,omitempty"`
	Nome         string     `json:"nome"`
	Idade        int        `json:"idade"`
}

func NewGetPessoaResponse(id uuid.UUID, createdAt time.Time, lastUpdateAt *time.Time, nome string, idade int) *GetPessoaResponse {
	return &GetPessoaResponse{
		ID:           id,
		CreatedAt:    createdAt,
		LastUpdateAt: lastUpdateAt,
		Nome:         nome,
		Idade:        idade,
	}
}
