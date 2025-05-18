package core

type DomainError struct {
	ErrorCode  int
	Message    string
	Err        string
	StatusCode int
}

func (e *DomainError) Error() string {
	return e.Message
}

type DomainErrorBuilder struct {
	DomainError DomainError
}

func NewDomainErrorBuilder(err error) *DomainErrorBuilder {
	return &DomainErrorBuilder{
		DomainError: DomainError{
			Err:        err.Error(),
			StatusCode: 400,
		},
	}
}

func (b *DomainErrorBuilder) WithErrorCode(errorCode int) *DomainErrorBuilder {
	b.DomainError.ErrorCode = errorCode
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
