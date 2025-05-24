package repositories

import "errors"

var (
	ErrUserNotFound      = errors.New("usuário não encontrado")
	ErrUserAlreadyExists = errors.New("usuário já existe")
)
