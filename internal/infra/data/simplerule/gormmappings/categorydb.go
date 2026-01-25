package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
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
	return db.Category
}

func (db *CategoryDB) FromDomain(d entities.Category) {
	db.Category = d
}

func SeedCategories(db *gorm.DB) {
	var count int64
	db.Model(&CategoryDB{}).Count(&count)
	if count > 0 {
		return // Já tem dados
	}

	categories := []CategoryDB{
		{Category: entities.NewCategory("Needs")},
		{Category: entities.NewCategory("Wants")},
		{Category: entities.NewCategory("Savings")},
	}

	for _, cat := range categories {
		db.Create(&cat)
	}
}
