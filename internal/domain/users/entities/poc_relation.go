package entities

import (
	"github.com/rkaturabara/flickly/internal/domain/core"
)

type PocRelation struct {
	core.Entity
	Nome              string `json:"nome"`
	PocRelationStatus string `json:"status"`
}
