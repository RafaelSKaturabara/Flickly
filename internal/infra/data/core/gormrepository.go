package core

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 2. O REPOSITÓRIO
type GormRepository[T any, M any] struct {
	db *gorm.DB
}

func NewGormRepository[T any, M any](db *gorm.DB) *GormRepository[T, M] {
	return &GormRepository[T, M]{db: db}
}

func (r *GormRepository[T, M]) GetByID(id uuid.UUID) (T, error) {
	var model M
	var domain T

	// Busca a struct de INFRA (M)
	err := r.db.First(&model, "id = ?", id).Error
	if err != nil {
		return domain, err
	}

	// Faz o CAST para a interface que criamos acima
	// any(&model) transforma em interface vazia para podermos testar o tipo
	if m, ok := any(&model).(PersistentEntity[T]); ok {
		return m.ToDomain(), nil
	}

	return domain, nil 
}

func (r *GormRepository[T, M]) Find(query func(T) bool) ([]T, error) {
	var allModels []M
	if err := r.db.Find(&allModels).Error; err != nil {
		return nil, err
	}

	var results []T
	for _, model := range allModels {
		var domain T
		if m, ok := any(&model).(PersistentEntity[T]); ok {
			domain = m.ToDomain()
		}

		if query(domain) {
			results = append(results, domain)
		}
	}
	return results, nil
}

func (r *GormRepository[T, M]) Create(entity T) error {
	var model M
	
	// Converte Domínio (T) para Infra (M) antes de salvar
	if m, ok := any(&model).(PersistentEntity[T]); ok {
		m.FromDomain(entity)
		return r.db.Save(m).Error
	}
	
	return r.db.Save(&entity).Error
}

func (r *GormRepository[T, M]) Delete(id uuid.UUID) error {
	var model M
	// No Delete, basta saber qual é a struct de Infra (M) para o GORM saber a tabela
	return r.db.Delete(&model, "id = ?", id).Error
}