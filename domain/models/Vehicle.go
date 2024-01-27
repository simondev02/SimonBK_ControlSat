package models

import (
	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	Plate       *string `gorm:"type:varchar(20):unique"`
	FkCompany   *int
	FkCustomer  *int
	FkAvlDevice *uint     `gorm:"type:varchar(20):unique"`
	AvlDevice   AvlDevice `gorm:"foreignKey:FkAvlDevice"`
	Company     Company   `gorm:"foreignKey:FkCompany"`
	Customer    Customer  `gorm:"foreignKey:FkCustomer"`
}
