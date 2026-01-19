package core

import (
	"log"
	"time"

	"gorm.io/driver/postgres" 
	"github.com/glebarez/sqlite" 
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetPostgresDBConnection() *gorm.DB {
	// 1. Configuração do DSN (Data Source Name)
	// Para SQLite, é apenas o nome do arquivo. 
	dsn := "simplerule.db"

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

func GetLocalDBConnection() *gorm.DB {
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