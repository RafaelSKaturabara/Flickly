package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type SimpleRuleUser struct {
	core.BaseEntity
	Nickname string 
}

func (u *SimpleRuleUser) IsValid() bool {
	return true
}
