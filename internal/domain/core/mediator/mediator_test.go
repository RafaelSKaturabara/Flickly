package mediator

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockRequest implementa a interface Request para testes
type MockRequest struct {
	Request
	Data string
}

// MockResponse implementa a interface Response para testes
type MockResponse struct {
	Result string
}

// MockHandler implementa a interface Handler para testes
type MockHandler struct {
	ReturnResponse MockResponse
	ReturnError    error
}

func (h *MockHandler) Handle(ctx context.Context, request Request) (Response, error) {
	return h.ReturnResponse, h.ReturnError
}

func TestNewMediatR(t *testing.T) {
	mediator := NewMediatR()
	assert.NotNil(t, mediator, "NewMediatR deve retornar uma instância não nula")
}

func TestRegisterAndSend(t *testing.T) {
	// Configuração
	mediator := NewMediatR()
	mockHandler := &MockHandler{
		ReturnResponse: MockResponse{Result: "success"},
		ReturnError:    nil,
	}
	mockRequest := MockRequest{Data: "test"}

	// Testar registro
	mediator.Register("MockRequest", mockHandler)

	// Testar envio de solicitação
	ctx := context.Background()
	response, err := mediator.Send(ctx, mockRequest)

	// Verificações
	assert.NoError(t, err, "Send não deve retornar erro quando o handler registrado retorna sucesso")

	mockResponse, ok := response.(MockResponse)
	assert.True(t, ok, "Response deve ser do tipo MockResponse")
	assert.Equal(t, "success", mockResponse.Result, "Response deve conter o resultado esperado")
}

func TestSendWithHandlerError(t *testing.T) {
	// Configuração
	mediator := NewMediatR()
	expectedError := errors.New("handler error")
	mockHandler := &MockHandler{
		ReturnResponse: MockResponse{},
		ReturnError:    expectedError,
	}
	mockRequest := MockRequest{Data: "test"}

	// Testar registro
	mediator.Register("MockRequest", mockHandler)

	// Testar envio de solicitação
	ctx := context.Background()
	response, err := mediator.Send(ctx, mockRequest)

	// Verificações
	assert.Equal(t, expectedError, err, "Send deve retornar o erro do handler")
	assert.Equal(t, MockResponse{}, response, "Response deve ser a resposta vazia do handler")
}

func TestSendWithoutRegisteredHandler(t *testing.T) {
	// Configuração
	mediator := NewMediatR()
	mockRequest := MockRequest{Data: "test"}

	// Testar envio de solicitação sem handler registrado
	ctx := context.Background()
	response, err := mediator.Send(ctx, mockRequest)

	// Verificações
	assert.Error(t, err, "Send deve retornar erro quando não há handler registrado")
	assert.Nil(t, response, "Response deve ser nil quando não há handler registrado")
	assert.Equal(t, "no handler registered for request type", err.Error(), "Mensagem de erro incorreta")
}
