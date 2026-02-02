package valueobjects

import (
	"database/sql/driver"
	"errors"

	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
)

// PasswordHash é um Value Object que garante a segurança da senha
type PasswordHash struct {
	hash string
}

// NewPasswordHash cria uma nova senha fazendo o hash automaticamente
func NewPasswordHash(plainText string) (PasswordHash, error) {
	if len(plainText) < 8 {
		return PasswordHash{}, errors.New("a senha deve ter pelo menos 8 caracteres")
	}

	hash, err := utilities.Encrypt(plainText)
	if err != nil {
		return PasswordHash{}, err
	}

	return PasswordHash{hash: string(hash)}, nil
}

// Compare verifica se a senha plana coincide com o hash
func (p PasswordHash) Compare(password string) bool {
	err := utilities.CompareHashAndPassword(p.hash, password)
	return err == nil
}

// GetHash retorna o hash da senha
func (p PasswordHash) GetHash() string {
	return p.hash
}

// Mapeamento do Value Object para o Banco de Dados
// Value: Transforma o objeto em algo que o Banco entende (String)
func (p PasswordHash) Value() (driver.Value, error) {
    return p.hash, nil
}

// Scan: Pega o que vem do Banco (String) e coloca no objeto
func (p *PasswordHash) Scan(value interface{}) error {
    if value == nil {
        return nil
    }
    s, ok := value.(string)
    if !ok {
        return errors.New("invalid data type for PasswordHash")
    }
    p.hash = s
    return nil
}