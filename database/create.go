package database

import (
	//--Entities Import--//
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(
	//--Entities Hook--//
	)

	if err != nil {
		return nil
	}

	return db
}
