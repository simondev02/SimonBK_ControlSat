package models

import (
	"gorm.io/gorm"
)

type AvlDevice struct {
	gorm.Model
	IMEI *string `gorm:"unique"`
}
