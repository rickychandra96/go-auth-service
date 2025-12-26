package database

import (
	"auth-service/internal/config"
	"auth-service/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.Database.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db, err
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.User{}, &domain.RefreshToken{})
	if err != nil {
		return err
	}
	return nil
}
