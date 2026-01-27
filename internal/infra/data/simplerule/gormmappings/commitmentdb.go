package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CommitmentDB struct {
	entities.Commitment `gorm:"embedded"`
	Description         string           `gorm:"type:varchar(255);not null"`
	UserSimpleRuleID              uuid.UUID        `gorm:"column:user_id;not null"`
	User                UserSimpleRuleDB `gorm:"foreignKey:UserID;references:ID"`
}

func (CommitmentDB) TableName() string {
	return "commitments"
}

func (db *CommitmentDB) ToDomain() entities.Commitment {
	var d entities.Commitment
	_ = copier.Copy(&d, db)
	return d
}

func (db *CommitmentDB) FromDomain(d entities.Commitment) {
	_ = copier.Copy(db, &d)	
}

func SeedCommitment(db *gorm.DB) {}
