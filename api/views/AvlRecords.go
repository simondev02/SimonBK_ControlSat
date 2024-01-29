package views

import "time"

type AvlRecord struct {
	Plate       *string    `json:"plate"`
	Imei        *string    `json:"imei"`
	Description *string    `json:"description"`
	Timestamp   *time.Time `json:"time_stamp_event"`
	Location    *string    `json:"location"`
	Latitude    *float64   `json:"latitude"`
	Longitude   *float64   `json:"longitude"`
	Altitude    *int       `json:"altitude"`
	Angle       *int       `json:"angle"`
	Satellites  *int       `json:"satellites"`
	Speed       *int       `json:"speed"`
	Event       *string    `json:"event"`
}
