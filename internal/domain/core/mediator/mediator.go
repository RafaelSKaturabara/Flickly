package mediator

import (
	"errors"
	"flickly/internal/infra/crosscutting/utilities"
	"github.com/gin-gonic/gin"
)

// Request interface para requisições
type Request interface {
}

// Response interface para respostas
type Response interface{}

// Handler interface para manipuladores de requisições
type Handler interface {
	Handle(c *gin.Context, request Request) (Response, error)
}

// Mediator interface para o mediador
type Mediator interface {
	Register(requestName string, handler Handler)
	Send(c *gin.Context, request Request) (Response, error)
}

type MediatR struct {
	handlers map[string]Handler
}

// NewMediator cria uma nova instância do MediatR
func NewMediatR() Mediator {
	return &MediatR{
		handlers: make(map[string]Handler),
	}
}

// Register registra um manipulador para um tipo de requisição
func (m *MediatR) Register(requestName string, handler Handler) {
	m.handlers[requestName] = handler
}

// Send envia a requisição para o manipulador apropriado
func (m *MediatR) Send(c *gin.Context, request Request) (Response, error) {
	structName := utilities.GetStructName(request)
	handler, ok := m.handlers[structName]
	if !ok {
		return nil, errors.New("no handler registered for request type")
	}
	return handler.Handle(c, request)
}
