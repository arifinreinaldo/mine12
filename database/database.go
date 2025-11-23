package database

import (
	"fmt"
	"log"

	"github.com/arifinreinaldo/mine12/config"
	"github.com/arifinreinaldo/mine12/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	var err error

	DB, err = gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established")
	return nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Customer{},
		&models.Sale{},
		&models.SaleItem{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
