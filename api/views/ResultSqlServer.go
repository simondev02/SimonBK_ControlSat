package views

import "time"

type ResultSqlServer struct {
	Imei        string    `json:"Imei"`
	Plate       string    `json:"Plate"`
	Description string    `json:"Description"`
	Speed       int16     `json:"Speed"`
	Latitude    float32   `json:"Latitude"`
	Longitude   float32   `json:"Longitude"`
	Timestamp   time.Time `json:"TimeStamp"`
	Event       string    `json:"Event"`
	Odometer    *float32  `json:"Odometer"`
}
