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
		{Subcategory: entities.Subcategory{Name: "Shelter", Description: "Rent, Mortgage, Property Taxes"}, Category: needs},
		{Subcategory: entities.Subcategory{Name: "Food", Description: "Groceries"}, Category: needs},
		{Subcategory: entities.Subcategory{Name: "Utilities", Description: "Water, Electricity, Internet"}, Category: needs},
		// Wants
		{Subcategory: entities.Subcategory{Name: "Dining Out", Description: "Restaurants, Coffee, Delivery"}, Category: wants},
		{Subcategory: entities.Subcategory{Name: "Entertainment", Description: "Movies, Concerts, Games"}, Category: wants},
		// Savings
		{Subcategory: entities.Subcategory{Name: "Emergency Fund", Description: "Safety net for 6 months"}, Category: savings},
		{Subcategory: entities.Subcategory{Name: "Investment", Description: "Stocks, Bonds, etc."}, Category: savings},
	}

	for _, sc := range subcats {
		db.Create(&sc)
	}
}
