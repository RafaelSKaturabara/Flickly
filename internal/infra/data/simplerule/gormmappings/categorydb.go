package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CategoryDB struct {
	entities.Category `gorm:"embedded"`
	Name              string `gorm:"type:varchar(255);not null"`
}

func (CategoryDB) TableName() string {
	return "categories"
}

func (db *CategoryDB) ToDomain() entities.Category {
	var d entities.Category
	_ = copier.Copy(&d, db)
	return d
}

func (db *CategoryDB) FromDomain(d entities.Category) {
	_ = copier.Copy(db, &d)
}

func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&CategoryDB{}).Count(&count)
	if count > 0 {
		return // Já tem dados
	}

	var needs, wants, savings entities.Category
	var needsDB, wantsDB, savingsDB CategoryDB

	needs = entities.NewCategory("Needs")
	wants = entities.NewCategory("Wants")
	savings = entities.NewCategory("Savings")

	needsDB.FromDomain(needs)
	wantsDB.FromDomain(wants)
	savingsDB.FromDomain(savings)

	categories := []CategoryDB{
		needsDB,
		wantsDB,
		savingsDB,
	}

	for _, cat := range categories {
		db.Create(&cat)
	}
}
