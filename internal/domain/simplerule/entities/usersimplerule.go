package entities

import "github.com/RafaelSKaturabara/Flickly/internal/domain/core"

type UserSimpleRule struct {
	core.BaseEntity
	Nickname string
}

func NewUserSimpleRule(nickname string) UserSimpleRule {
	return UserSimpleRule{
		BaseEntity: core.NewBaseEntity(),
		Nickname:   nickname,
	}
}

func (u *UserSimpleRule) IsValid() bool {
	return true
}
