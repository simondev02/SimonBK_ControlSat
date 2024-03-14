package views

import "time"

type AvlRecord struct {
	Plate          *string    `json:"plate"`
	TimeStampEvent *time.Time `json:"time_stamp_event"`
	Location       *string    `json:"location"`
	Latitude       *float64   `json:"latitude"`
	Longitude      *float64   `json:"longitude"`
	Angle          *int       `json:"angle"`
	Speed          *int16     `json:"speed"`
	Odometer       *float32   `json:"odometer"`
}
