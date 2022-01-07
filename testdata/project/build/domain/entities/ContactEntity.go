package entities

import "gorm.io/gorm"

type ContactEntity struct {
	gorm.Model
	Uuid    string
	Name    string
	Phone   string
	Age     int
	Enabled bool
}

func NewContactEntity(uuid string, name string, phone string, age int, enabled bool) *ContactEntity {
	return &ContactEntity{
		Uuid:    uuid,
		Name:    name,
		Phone:   phone,
		Age:     age,
		Enabled: enabled,
	}
}
