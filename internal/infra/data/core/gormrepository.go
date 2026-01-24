package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) GetByID(id uuid.UUID) (T, error) {
	var entity T
	// O banco cuida do bloqueio/isolamento automaticamente
	err := r.db.First(&entity, "id = ?", id).Error
	return entity, err
}

func (r *GormRepository[T]) Find(query func(T) bool) ([]T, error) {
	var allItems []T
	if err := r.db.Find(&allItems).Error; err != nil {
		return nil, err
	}

	// Como sua interface usa uma func(T) bool, ainda filtramos em memória após a busca
	var results []T
	for _, item := range allItems {
		if query(item) {
			results = append(results, item)
		}
	}
	return results, nil
}

func (r *GormRepository[T]) Create(entity T) error {
	// O GORM gerencia a transação de escrita (Write)
	return r.db.Save(&entity).Error
}

func (r *GormRepository[T]) Delete(id uuid.UUID) error {
	var entity T
	return r.db.Delete(&entity, "id = ?", id).Error
}
