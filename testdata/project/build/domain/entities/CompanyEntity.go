package entities

import "gorm.io/gorm"

type CompanyEntity struct {
	gorm.Model
	Uuid    string
	Name    string
	Phone   string
	Address string
	Enabled bool
}

func NewCompanyEntity(uuid string, name string, phone string, address string, enabled bool) *CompanyEntity {
	return &CompanyEntity{
		Uuid:    uuid,
		Name:    name,
		Phone:   phone,
		Address: address,
		Enabled: enabled,
	}
}
