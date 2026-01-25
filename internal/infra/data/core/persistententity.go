package core

// Todo modelo de INFRA (DB) deve implementar isso
type PersistentEntity[T any] interface {
    ToDomain() T      // Converte DB -> Domínio
    FromDomain(T)     // Converte Domínio -> DB
}