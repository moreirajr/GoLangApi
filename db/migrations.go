package db

import (
	"dependencies/entities"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) error {
	return db.AutoMigrate(&entities.Employee{})
}
