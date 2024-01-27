package views

import "time"

type ResultAll struct {
	Imei       string    `json:"Imei"`
	Plate      string    `json:"Plate"`
	Latitude   float32   `json:"Latitude"`
	Longitude  float32   `json:"Longitude"`
	Timestamp  time.Time `json:"TimeStamp"`
	FkCompany  *int      `json:"FkCompany"`
	Company    *string   `json:"Company"`
	FkCustomer *int      `json:"FkCustomer"`
	Customer   *string   `json:"Customer"`
	State      *string   `json:"State"`
}
