package core

import "errors"

type DomainError struct {
	error
	Code       int
	Message    string
	StatusCode int
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

func (b *DomainErrorBuilder) Build() *DomainError {
	return &b.DomainError
}

var (
	ErrUserAlreadyExist = func(err error) *DomainError {
		return NewDomainErrorBuilder(err).WithMessage("Usuário já cadastrado").WithErrorCode(1).Build()
	}
)
