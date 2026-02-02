package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type Client struct {
	core.BaseEntity
	Secret       string   // A "senha" do aplicativo
	Name         string   // Nome do App (ex: "App Financeiro")
	RedirectURIs []RedirectURI // Segurança: onde o OAuth pode devolver o código
}

func NewClientWithRedirectURIs(name, secret string, redirectURIs []RedirectURI) Client {
	return Client{
		BaseEntity: core.NewBaseEntity(),
		Secret:     secret,
		Name:       name,
		RedirectURIs: redirectURIs,
	}
}

func NewClient(name, secret string) Client {
	return Client{
		BaseEntity: core.NewBaseEntity(),
		Secret:     secret,
		Name:       name,
	}
}

func (c *Client) IsValid() bool {
	return true
}
