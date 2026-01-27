package entities

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
)

type PocRelation struct {
	core.Entity
	Nome              string `json:"nome"`
	PocRelationStatus string `json:"status"`
}
