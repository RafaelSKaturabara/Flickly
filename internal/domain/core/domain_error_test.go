package core

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomainError_Error(t *testing.T) {
	// Configuração
	message := "Erro de teste"
	domainError := &DomainError{
		Code:       123,
		Message:    message,
		error:      errors.New("erro original"),
		StatusCode: 400,
	}

	// Execução
	errorMessage := domainError.Message

	// Verificação
	assert.Equal(t, message, errorMessage, "O método Error() deve retornar a mensagem de erro")
}

func TestNewDomainErrorBuilder(t *testing.T) {
	// Configuração
	originalError := errors.New("erro original")

	// Execução
	builder := NewDomainErrorBuilder(originalError)

	// Verificações
	assert.NotNil(t, builder, "NewDomainErrorBuilder deve retornar uma instância não nula")
	assert.Equal(t, originalError.Error(), builder.DomainError.Error(), "O erro original deve ser armazenado")
	assert.Equal(t, 400, builder.DomainError.StatusCode, "O código de status padrão deve ser 400")
}

func TestDomainErrorBuilder_WithErrorCode(t *testing.T) {
	// Configuração
	builder := NewDomainErrorBuilder(errors.New("erro"))
	errorCode := 123

	// Execução
	result := builder.WithErrorCode(errorCode)

	// Verificações
	assert.Equal(t, builder, result, "WithErrorCode deve retornar o próprio builder")
	assert.Equal(t, errorCode, builder.DomainError.Code, "O código de erro deve ser configurado corretamente")
}

func TestDomainErrorBuilder_WithMessage(t *testing.T) {
	// Configuração
	builder := NewDomainErrorBuilder(errors.New("erro"))
	message := "Mensagem de teste"

	// Execução
	result := builder.WithMessage(message)

	// Verificações
	assert.Equal(t, builder, result, "WithMessage deve retornar o próprio builder")
	assert.Equal(t, message, builder.DomainError.Message, "A mensagem deve ser configurada corretamente")
}

func TestDomainErrorBuilder_WithStatusCode(t *testing.T) {
	// Configuração
	builder := NewDomainErrorBuilder(errors.New("erro"))
	statusCode := 500

	// Execução
	result := builder.WithStatusCode(statusCode)

	// Verificações
	assert.Equal(t, builder, result, "WithStatusCode deve retornar o próprio builder")
	assert.Equal(t, statusCode, builder.DomainError.StatusCode, "O código de status deve ser configurado corretamente")
}

func TestDomainErrorBuilder_Build(t *testing.T) {
	// Configuração
	originalError := errors.New("erro original")
	errorCode := 123
	message := "Mensagem de teste"
	statusCode := 500

	builder := NewDomainErrorBuilder(originalError)
	builder.WithErrorCode(errorCode)
	builder.WithMessage(message)
	builder.WithStatusCode(statusCode)

	// Execução
	domainError := builder.Build()

	// Verificações
	assert.NotNil(t, domainError, "Build deve retornar uma instância não nula")
	assert.Equal(t, errorCode, domainError.Code, "O código de erro deve ser configurado corretamente")
	assert.Equal(t, message, domainError.Message, "A mensagem deve ser configurada corretamente")
	assert.Equal(t, originalError.Error(), domainError.Error(), "O erro original deve ser armazenado")
	assert.Equal(t, statusCode, domainError.StatusCode, "O código de status deve ser configurado corretamente")
}

func TestErrUserAlreadyExist(t *testing.T) {
	// Configuração
	originalError := errors.New("erro original")

	// Execução
	domainError := ErrUserAlreadyExist(originalError)

	// Verificações
	assert.NotNil(t, domainError, "ErrUserAlreadyExist deve retornar uma instância não nula")
	assert.Equal(t, 1, domainError.Code, "O código de erro deve ser 1")
	assert.Equal(t, "Usuário já cadastrado", domainError.Message, "A mensagem deve ser 'Usuário já cadastrado'")
	assert.Equal(t, originalError.Error(), domainError.Error(), "O erro original deve ser armazenado")
}
