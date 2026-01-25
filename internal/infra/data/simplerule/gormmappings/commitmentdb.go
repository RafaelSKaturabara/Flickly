package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommitmentDB struct {
	entities.Commitment `gorm:"embedded"`
	Description         string           `gorm:"type:varchar(255);not null"`
	UserID              uuid.UUID        `gorm:"column:user_id;not null"`
	User                SimpleRuleUserDB `gorm:"foreignKey:UserID;references:ID"`
}

func (CommitmentDB) TableName() string {
	return "commitments"
}

func (db *CommitmentDB) ToDomain() entities.Commitment {
	return db.Commitment
}

func (db *CommitmentDB) FromDomain(d entities.Commitment) {
	db.Commitment = d
}

func SeedCommitment(db *gorm.DB) {}
