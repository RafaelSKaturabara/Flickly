package auth

import (
	"fmt"
	"os"
)

type AuthConfig struct {
	JWTSecret    string
	ClientID     string
	ClientSecret string
}

func LoadAuthConfig() (*AuthConfig, error) {
	config := &AuthConfig{
		JWTSecret:    os.Getenv("JWT_SECRET"),
		ClientID:     os.Getenv("OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH2_CLIENT_SECRET"),
	}

	// Validação das variáveis obrigatórias
	if config.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET não configurado")
	}
	if config.ClientID == "" {
		return nil, fmt.Errorf("OAUTH2_CLIENT_ID não configurado")
	}
	if config.ClientSecret == "" {
		return nil, fmt.Errorf("OAUTH2_CLIENT_SECRET não configurado")
	}

	return config, nil
}
