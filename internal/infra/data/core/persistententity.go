package core

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
)

// Todo modelo de INFRA (DB) deve implementar isso
type PersistentEntity[T any] interface {
	ToDomain() T  // Converte DB -> Domínio
	FromDomain(T) // Converte Domínio -> DB
}

// BaseEntityDB fornece o mapeamento padrão do GORM para os campos do BaseEntity do domínio
type BaseEntityDB struct {
	core.BaseEntity `gorm:"embedded"`
}

func (db *BaseEntityDB) FromBaseDomain(d core.BaseEntity) {
	db.BaseEntity = d
}

func (db *BaseEntityDB) ToBaseDomain() core.BaseEntity {
	return db.BaseEntity
}
