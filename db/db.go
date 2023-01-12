package db

import (
	"dependencies/config"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg *config.Config) *gorm.DB {
	log.Printf("\n\n[Trying to connect on database %s]\n\n", cfg.ConnectionString)
	db, err := gorm.Open(sqlserver.Open(cfg.ConnectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
