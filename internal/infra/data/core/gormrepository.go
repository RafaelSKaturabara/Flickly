package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 2. O REPOSITÓRIO
type GormRepository[T any, M any] struct {
	DB *gorm.DB
}

func NewGormRepository[T any, M any](db *gorm.DB) *GormRepository[T, M] {
	return &GormRepository[T, M]{DB: db}
}

func (r *GormRepository[T, M]) GetByID(id uuid.UUID) (*T, error) {
	var model M

	// Busca a struct de INFRA (M)
	err := r.DB.First(&model, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	// Faz o CAST para a interface que criamos acima
	// any(&model) transforma em interface vazia para podermos testar o tipo
	if m, ok := any(&model).(PersistentEntity[T]); ok {
		domain := m.ToDomain()
		return &domain, nil
	}

	return nil, nil
}

func (r *GormRepository[T, M]) Find(query func(*T) bool) ([]*T, error) {
	var allModels []M
	if err := r.DB.Find(&allModels).Error; err != nil {
		return nil, err
	}

	var results []*T
	for _, model := range allModels {
		if m, ok := any(&model).(PersistentEntity[T]); ok {
			domain := m.ToDomain()
			if query(&domain) {
				results = append(results, &domain)
			}
		}
	}
	return results, nil
}

func (r *GormRepository[T, M]) Create(entity *T) error {
	var model M

	// Converte Domínio (T) para Infra (M) antes de salvar
	if m, ok := any(&model).(PersistentEntity[T]); ok {
		m.FromDomain(*entity)
		return r.DB.Save(m).Error
	}

	return r.DB.Save(entity).Error
}

func (r *GormRepository[T, M]) Update(entity *T) error {
	var model M

	// Converte Domínio (T) para Infra (M) antes de salvar
	if m, ok := any(&model).(PersistentEntity[T]); ok {
		m.FromDomain(*entity)
		return r.DB.Updates(m).Error
	}

	return r.DB.Updates(entity).Error
}

func (r *GormRepository[T, M]) Delete(id uuid.UUID) error {
	var model M
	// No Delete, basta saber qual é a struct de Infra (M) para o GORM saber a tabela
	return r.DB.Delete(&model, "id = ?", id).Error
}
