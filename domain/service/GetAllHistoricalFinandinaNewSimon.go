package service

import (
	"SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/domain/models"

	"gorm.io/gorm"
)

func GetLastRecordForEachPlate(db *gorm.DB) ([]views.AvlRecord, error) {
	var records []models.AvlRecord
	var responseRecords []views.AvlRecord

	// Obtener todos los registros donde Id_company sea 12
	err := db.Where("id_company = ?", 12).Find(&records).Error
	if err != nil {
		return nil, err
	}

	// Crear un mapa para almacenar el último registro de cada plate
	lastRecordForEachPlate := make(map[string]models.AvlRecord)

	// Iterar sobre los registros y almacenar el último registro de cada plate
	for _, record := range records {
		if lastRecord, ok := lastRecordForEachPlate[*record.Plate]; ok {
			// Si el registro actual es más reciente que el último registro almacenado, actualizar el último registro
			if record.TimeStampEvent.After(*lastRecord.TimeStampEvent) {
				lastRecordForEachPlate[*record.Plate] = record
			}
		} else {
			// Si no hay ningún registro almacenado para esta plate, almacenar el registro actual
			lastRecordForEachPlate[*record.Plate] = record
		}
	}

	// Iterar sobre el último registro de cada plate y mapearlos a responseRecords
	for _, record := range lastRecordForEachPlate {
		responseRecord := views.AvlRecord{
			Plate:      record.Plate,
			Imei:       record.Imei,
			Timestamp:  record.TimeStampEvent,
			Location:   record.Location,
			Latitude:   record.Latitude,
			Longitude:  record.Longitude,
			Altitude:   record.Altitude,
			Angle:      record.Angle,
			Satellites: record.Satellites,
			Speed:      record.Speed,
			Event:      record.Event,
		}
		responseRecords = append(responseRecords, responseRecord)
	}

	return responseRecords, nil
}
