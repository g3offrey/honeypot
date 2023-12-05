package persistence

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("can't init db: %w", err)
	}

	err = db.AutoMigrate(&CatchModel{})
	if err != nil {
		return nil, fmt.Errorf("can't migrate db: %w", err)
	}

	return db, nil
}
