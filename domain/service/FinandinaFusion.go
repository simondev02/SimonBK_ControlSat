package service

import (
	"SimonBK_ControlSat/api/views"

	"gorm.io/gorm"
)

func GetAllRecords(db *gorm.DB) ([]views.AvlRecord, error) {
	// Llamar a GetAllFinandina
	allFinandina, err := GetAllFinandina()
	if err != nil {
		return nil, err
	}

	// Convertir allFinandina a []views.AvlRecord
	var allFinandinaAvl []views.AvlRecord
	for _, record := range allFinandina {
		latitude := float64(*record.Latitude)
		longitude := float64(*record.Longitude)
		allFinandinaAvl = append(allFinandinaAvl, views.AvlRecord{
			Plate:       record.Plate,
			Imei:        record.Imei,
			Description: record.Description,
			Timestamp:   record.Timestamp,
			Latitude:    &latitude,
			Longitude:   &longitude,
			Event:       record.Event,
		})
	}

	// Llamar a GetLastRecordForEachPlate
	lastRecords, err := GetLastRecordForEachPlate(db)
	if err != nil {
		return nil, err
	}

	// Unir los resultados
	allRecords := append(allFinandinaAvl, lastRecords...)

	return allRecords, nil
}
