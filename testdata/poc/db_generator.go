package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	openDatabase()
}

type Entity struct {
	gorm.Model
	Name      string     `json:"name"`
	Fields    []Field    `json:"fields"`
	Relations []Relation `json:"relations"`
}

type Field struct {
	gorm.Model
	EntityID int
	Name     string `json:"name"`
	Type     string `json:"type"`
	Length   int    `json:"length"`
}

type Relation struct {
	gorm.Model
	EntityID int
	Name     string `json:"name"`
	Type     string `json:"type"`
	Relation string `json:"Relation"`
}

func openDatabase() *gorm.DB {
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Entity{}, &Field{}, &Relation{})

	// Create
	db.Create(&Entity{Name: "Customer", Fields: []Field{{Name: "namEs", Type: "xtring"}}})
	db.Create(&Entity{Name: "Product", Fields: []Field{{Name: "concept", Type: "string"}}})
	db.Create(&Entity{Name: "Provider", Fields: []Field{{Name: "price", Type: "int"}}})

	// Read
	var field Field
	db.First(&field, 1)                   // find product with integer primary key
	db.First(&field, "name = ?", "namEs") // find product with code D42

	// Update - update product's price to 200
	db.Model(&field).Update("Name", "NAme")
	// Update - update multiple fields
	db.Model(&field).Updates(Field{Name: "names", Type: "String"}) // non-zero fields
	db.Model(&field).Updates(map[string]interface{}{"Name": "name", "type": "string"})

	// Delete - delete product
	db.Delete(&field, 1)
	var product Entity
	db.First(&product, Entity{Name: "Provider"})
	db.Delete(&product.Fields, &product.ID)
	db.Delete(&product.Relations, &product.ID)
	db.Delete(&product)

	// Get all records
	var results []Entity
	db.Find(&results)
	fmt.Printf("---- %s ----", "word")
	println("Entities in DB:\n")
	for _, record := range results {
		println(record.Name)
	}

	return db
}
