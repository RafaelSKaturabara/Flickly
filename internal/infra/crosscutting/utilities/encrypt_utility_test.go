package utilities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Senha válida",
			password: "senha123",
			wantErr:  false,
		},
		{
			name:     "Senha vazia",
			password: "",
			wantErr:  false,
		},
		{
			name:     "Senha longa",
			password: "senha_muito_longa_com_muitos_caracteres_123456789",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Execução
			hashedPassword, err := Encrypt(tt.password)

			// Verificações
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, hashedPassword)
			assert.NotEqual(t, tt.password, hashedPassword, "A senha não deve ser armazenada em texto puro")
		})
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		getHash  func() string
		wantErr  bool
	}{
		{
			name:     "Senha correta",
			password: "senha123",
			getHash: func() string {
				hash, _ := Encrypt("senha123")
				return hash
			},
			wantErr: false,
		},
		{
			name:     "Senha incorreta",
			password: "senha_errada",
			getHash: func() string {
				hash, _ := Encrypt("senha123")
				return hash
			},
			wantErr: true,
		},
		{
			name:     "Hash inválido",
			password: "senha123",
			getHash: func() string {
				return "hash_invalido"
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword := tt.getHash()
			err := CompareHashAndPassword(hashedPassword, tt.password)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestEncryptAndCompare(t *testing.T) {
	// Configuração
	password := "senha123"

	// Execução
	hashedPassword, err := Encrypt(password)
	assert.NoError(t, err)

	// Verificações
	err = CompareHashAndPassword(hashedPassword, password)
	assert.NoError(t, err, "A senha deve ser verificada corretamente após a criptografia")

	// Teste com senha incorreta
	err = CompareHashAndPassword(hashedPassword, "senha_errada")
	assert.Error(t, err, "Deve retornar erro ao comparar com senha incorreta")
}
