package gormmappings

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
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
	var d entities.Subcategory
	copier.Copy(&d, db)
	return d
}

func (db *SubcategoryDB) FromDomain(d entities.Subcategory) {
	copier.Copy(db, &d)
}

func SeedSubcategories(db *gorm.DB) {
	var count int64
	db.Model(&SubcategoryDB{}).Count(&count)
	if count > 0 {
		return
	}

	// Precisamos buscar os IDs reais das categorias criadas no seed anterior
	var needsDB, wantsDB, savingsDB CategoryDB
	db.Where("name = ?", "Needs").First(&needsDB)
	db.Where("name = ?", "Wants").First(&wantsDB)
	db.Where("name = ?", "Savings").First(&savingsDB)

	var shelter, food, utilities, diningOut, entertainment, emergencyFund, investment entities.Subcategory

	var shelterDB, foodDB, utilitiesDB, diningOutDB, entertainmentDB, emergencyFundDB, investmentDB SubcategoryDB

	shelter = entities.NewSubcategory("Shelter", "Rent, Mortgage, Property Taxes", needsDB.ID)
	food = entities.NewSubcategory("Food", "Groceries", needsDB.ID)
	utilities = entities.NewSubcategory("Utilities", "Water, Electricity, Internet", needsDB.ID)
	diningOut = entities.NewSubcategory("Dining Out", "Restaurants, Coffee, Delivery", wantsDB.ID)
	entertainment = entities.NewSubcategory("Entertainment", "Movies, Concerts, Games", wantsDB.ID)
	emergencyFund = entities.NewSubcategory("Emergency Fund", "Safety net for 6 months", savingsDB.ID)
	investment = entities.NewSubcategory("Investment", "Stocks, Bonds, etc.", savingsDB.ID)

	shelterDB.FromDomain(shelter)
	foodDB.FromDomain(food)
	utilitiesDB.FromDomain(utilities)
	diningOutDB.FromDomain(diningOut)
	entertainmentDB.FromDomain(entertainment)
	emergencyFundDB.FromDomain(emergencyFund)
	investmentDB.FromDomain(investment)

	subcats := []SubcategoryDB{
		// Needs
		shelterDB,
		foodDB,
		utilitiesDB,
		// Wants
		diningOutDB,
		entertainmentDB,
		// Savings
		emergencyFundDB,
		investmentDB,
	}

	for _, sc := range subcats {
		db.Create(&sc)
	}
}
