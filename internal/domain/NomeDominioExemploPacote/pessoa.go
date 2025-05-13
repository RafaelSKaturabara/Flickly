package NomeDominioExemploPacote

import "flickly/internal/domain/core"

type Pessoa struct {
	core.Entity
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func NewPessoa(nome string, idade int) *Pessoa {
	return &Pessoa{
		Entity: core.NewEntity(),
		Nome:   nome,
		Idade:  idade,
	}
}
