package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type DatabaseConfig struct {
	Filename string
}

func InitDatabase(config *DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Filename), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database")
	}

	err = db.AutoMigrate(User{}, Activity{}, ActivityLog{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
