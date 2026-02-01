package users

import (
	"log"
	"time"

	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/gormmappings"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MigrateDatabase(db *gorm.DB) {
	log.Println("[Database] Iniciando migrações...")

	// 1. Rodar AutoMigrate
	err := db.AutoMigrate(
		&gormmappings.UserDB{},
		&gormmappings.ClientDB{},
		&gormmappings.RoleDB{},
		&gormmappings.RedirectURIDB{},
		&gormmappings.UserAccessDB{},
		&gormmappings.AccessGrantDB{},
	)
	if err != nil {
		log.Fatalf("[Database] Erro ao migrar: %v", err)
	}

	// 2. Rodar Seeds (Ordem importa por causa de FKs)
	log.Println("[Database] Verificando sementes (seeds)...")
	gormmappings.SeedAccessGrants(db)
	gormmappings.SeedClients(db)
	gormmappings.SeedRedirectURIs(db)
	gormmappings.SeedRoles(db)
	gormmappings.SeedUserAccess(db)
	gormmappings.SeedUsers(db)

	log.Println("[Database] Infraestrutura de dados pronta.")
}

func GetIdentityPostgresDBConnection() *gorm.DB {
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

func GetIdentityLocalDBConnection() *gorm.DB {
	// 1. Configuração do DSN (Data Source Name)
	// Para SQLite, é apenas o nome do arquivo.
	dsn := "identity.db"

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
