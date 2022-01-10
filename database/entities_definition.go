package database

import "gorm.io/gorm"

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
}

type Relation struct {
	gorm.Model
	EntityID int
	Name     string `json:"name"`
	Type     string `json:"type"`
	Relation string `json:"Relation"`
}
