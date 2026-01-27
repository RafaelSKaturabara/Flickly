package core

import (
	"errors"
	"net/http"
)

type DomainError struct {
	error
	Code       int
	Message    string
	StatusCode int
	Errors    *map[string]string
}

type DomainErrorBuilder struct {
	DomainError DomainError
}

func NewDomainErrorBuilder(err error) *DomainErrorBuilder {
	defaultErr := errors.New("generic error")
	if err != nil {
		defaultErr = err
	}
	return &DomainErrorBuilder{
		DomainError: DomainError{
			error:      defaultErr,
			StatusCode: 400,
		},
	}
}

func (b *DomainErrorBuilder) WithErrorCode(code int) *DomainErrorBuilder {
	b.DomainError.Code = code
	return b
}

func (b *DomainErrorBuilder) WithMessage(message string) *DomainErrorBuilder {
	b.DomainError.Message = message
	return b
}

func (b *DomainErrorBuilder) WithStatusCode(statusCode int) *DomainErrorBuilder {
	b.DomainError.StatusCode = statusCode
	return b
}

func (b *DomainErrorBuilder) WithErrors(errors *map[string]string) *DomainErrorBuilder {
	b.DomainError.Errors = errors
	return b
}

func (b *DomainErrorBuilder) Build() *DomainError {
	return &b.DomainError
}

var (
	// Core
	DoNotUseThisGenericError = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("Erro genérico").WithErrorCode(0).Build()
	}
	InvalidEntityError = func(errors *map[string]string) *DomainError {
		return NewDomainErrorBuilder(nil).WithMessage("Entidade inválida").WithErrors(errors).WithErrorCode(1).Build()
	}

	// Users
	ErrUserAlreadyExist = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("Usuário já cadastrado").WithErrorCode(100).Build()
	}
	ErrInvalidCredentials = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("Credenciais inválidas").WithErrorCode(101).WithStatusCode(http.StatusUnauthorized).Build()
	}
	ErrUserNotFound = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("usuário não encontrado").WithErrorCode(102).Build()
	}
	ErrInvalidGrant = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("grant type inválido").WithErrorCode(103).Build()
	}
	ErrInvalidToken = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("Token inválido").WithErrorCode(104).WithStatusCode(http.StatusUnauthorized).Build()
	}
	ErrInvalidClient = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("client_id inválido").WithErrorCode(105).Build()
	}
	ErrInvalidScope = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("escopo inválido").WithErrorCode(106).Build()
	}

	// 
)
