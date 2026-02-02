package simplerule

import (
	"log"
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/gormmappings"
	usergormmappings "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MigrateSimpleRuleDatabase(db *gorm.DB) {
	log.Println("[Database] Iniciando migrações...")

	// 1. Rodar AutoMigrate
	err := db.AutoMigrate(
		&gormmappings.CategoryDB{},
		&gormmappings.SubcategoryDB{},
		&gormmappings.CommitmentDB{},
		&gormmappings.ExpenseDB{},
		&gormmappings.IncomeDB{},
		&gormmappings.UserSimpleRuleDB{},
		&usergormmappings.UserDB{},
	)
	if err != nil {
		log.Fatalf("[Database] Erro ao migrar: %v", err)
	}

	// 2. Rodar Seeds (Ordem importa por causa de FKs)
	log.Println("[Database] Verificando sementes (seeds)...")
	gormmappings.SeedCategories(db)
	gormmappings.SeedSubcategories(db)
	gormmappings.SeedCommitment(db)
	gormmappings.SeedExpense(db)
	gormmappings.SeedIncome(db)
	gormmappings.SeedUserSimpleRule(db)

	log.Println("[Database] Infraestrutura de dados pronta.")
}

func GetPostgresSimpleRuleDBConnection() *gorm.DB {
	// 1. Configuração do DSN (Data Source Name)
	// Exemplo: "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "host=localhost user=postgres password=postgres dbname=flickly port=5432 sslmode=disable"

	// 2. Abrir a conexão
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Loga todas as queries SQL (ótimo para debug)
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// 3. Configurações de Pool de Conexão (Opcional, mas recomendado)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

func GetLocalSimpleRuleDBConnection() *gorm.DB {
	// 1. Configuração do DSN (Data Source Name)
	// Para SQLite, é apenas o nome do arquivo.
	dsn := "simplerule.db"

	// 2. Abrir a conexão
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Loga todas as queries SQL (ótimo para debug)
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// 3. Configurações de Pool de Conexão (Opcional, mas recomendado)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
