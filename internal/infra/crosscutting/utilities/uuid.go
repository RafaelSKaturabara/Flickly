package utilities

import (
	"github.com/google/uuid"
)

// ParseUUID converte uma string em UUID
func ParseUUID(id string) uuid.UUID {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil
	}
	return parsedUUID
}
