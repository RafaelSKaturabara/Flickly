package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
)

type SimpleRuleUserDB struct {
	entities.SimpleRuleUser `gorm:"embedded"`
	Nickname                string `gorm:"type:varchar(100);not null"`
}

func (SimpleRuleUserDB) TableName() string {
	return "simplerule_users"
}

func (db *SimpleRuleUserDB) ToDomain() entities.SimpleRuleUser {
	return db.SimpleRuleUser
}

func (db *SimpleRuleUserDB) FromDomain(d entities.SimpleRuleUser) {
	db.SimpleRuleUser = d
}
