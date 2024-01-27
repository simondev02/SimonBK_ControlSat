package models

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name                 *string
	Address              *string
	PhoneNumber          *string
	Email                *string
	IdentificationNumber *string `gorm:"uniqueIndex:idx_fkcompany_identification"`
	DeletedBy            *uint
	UpdatedBy            *uint
	FkCompany            *int    `gorm:"uniqueIndex:idx_fkcompany_identification"`
	Company              Company `gorm:"foreignKey:FkCompany"`
}
