package views

import "time"

type AvlRecords struct {
	Plate          string `gorm:"type:varchar(10);not null"`
	Imei           string `gorm:"type:varchar(20);not null"`
	FkManufacturer int
	Manufacturer   string
	Vehicle_type   int
	Ip             string
	TimeStampEvent time.Time
	Id_company     int
	Company        string
	Id_customer    int
	Customer       string
	Location       string
	Latitude       float32 `gorm:"not null"`
	Longitude      float32 `gorm:"not null"`
	Altitude       int16
	Angle          int16
	Satellites     int8
	Speed          int16
	Hdop           int16
	Pdop           int16
	Event          string
	Is_alarm       int
	Properties     string `gorm:"column:properties"`
}
