package entities

import (
	"database/sql/driver"
	"fmt"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

// Valuer: Permite que o GORM grave o Enum no banco como string
func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

// Scanner: Permite que o GORM leia a string do banco e converta de volta para Role
func (r *Role) Scan(value interface{}) error {
	sv, ok := value.(string)
	if !ok {
		return fmt.Errorf("tipo inválido para Role: %T", value)
	}
	*r = Role(sv)
	return nil
}