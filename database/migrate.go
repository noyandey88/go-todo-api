package database

import "gorm.io/gorm"

func migrate(db *gorm.DB) error {
	return db.AutoMigrate() // TODO: models will implement here
}
