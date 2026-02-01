package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type Client struct {
	core.BaseEntity
	Secret       string   // A "senha" do aplicativo
	Name         string   // Nome do App (ex: "App Financeiro")
	RedirectURIs []string // Segurança: onde o OAuth pode devolver o código
}

func NewClient(name, secret string, redirectURIs []string) Client {
	return Client{
		BaseEntity: core.NewBaseEntity(),
		Secret:     secret,
		Name:       name,
		RedirectURIs: redirectURIs,
	}
}

func (c *Client) IsValid() bool {
	return true
}
