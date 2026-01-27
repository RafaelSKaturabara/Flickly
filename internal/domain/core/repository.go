package core

import "github.com/google/uuid"

type Repository[T any] interface {
	GetByID(id uuid.UUID) (*T, error)
	Find(query func(*T) bool) ([]*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uuid.UUID) error
}
