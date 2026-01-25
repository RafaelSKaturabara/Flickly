package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubcategoryDB struct {
	entities.Subcategory `gorm:"embedded"`
	ID                   uint       `gorm:"primaryKey"`
	Name                 string     `gorm:"type:varchar(255);not null"`
	CategoryID           uuid.UUID  `gorm:"column:category_id;not null"`
	Category             CategoryDB `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Relacionamento explícito
}

func (SubcategoryDB) TableName() string {
	return "subcategories"
}

func (db *SubcategoryDB) ToDomain() entities.Subcategory {
	return db.Subcategory
}

func (db *SubcategoryDB) FromDomain(d entities.Subcategory) {
	db.Subcategory = d
}

func SeedSubcategories(db *gorm.DB) {
	var count int64
	db.Model(&SubcategoryDB{}).Count(&count)
	if count > 0 {
		return
	}

	// Precisamos buscar os IDs reais das categorias criadas no seed anterior
	var needs, wants, savings CategoryDB
	db.Where("name = ?", "Needs").First(&needs)
	db.Where("name = ?", "Wants").First(&wants)
	db.Where("name = ?", "Savings").First(&savings)

	subcats := []SubcategoryDB{
		// Needs
		{Subcategory: entities.NewSubcategory("Shelter", "Rent, Mortgage, Property Taxes", needs.ID)},
		{Subcategory: entities.NewSubcategory("Food", "Groceries", needs.ID)},
		{Subcategory: entities.NewSubcategory("Utilities", "Water, Electricity, Internet", needs.ID)},
		// Wants
		{Subcategory: entities.NewSubcategory("Dining Out", "Restaurants, Coffee, Delivery", wants.ID)},
		{Subcategory: entities.NewSubcategory("Entertainment", "Movies, Concerts, Games", wants.ID)},
		// Savings
		{Subcategory: entities.NewSubcategory("Emergency Fund", "Safety net for 6 months", savings.ID)},
		{Subcategory: entities.NewSubcategory("Investment", "Stocks, Bonds, etc.", savings.ID)},
	}

	for _, sc := range subcats {
		db.Create(&sc)
	}
}
