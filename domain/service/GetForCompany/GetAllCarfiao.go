package service

import (
	"SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/infra/db"
	"encoding/json"
	"fmt"
)

func GetAllCarfiao() ([]views.AvlRecords, error) {
	results, err := GetAllFinandina(4)
	if err != nil {
		return nil, err
	}
	Id_company := 4
	Company := "Carfiao"

	avlRecords := make([]views.AvlRecords, len(results))
	for i, result := range results {
		properties := views.Properties{
			TotalOdometer: result.Odometer,
		}
		propertiesJson, err := json.Marshal(properties)
		if err != nil {
			return nil, fmt.Errorf("error al convertir Properties a JSON: %w", err)
		}
		avlRecords[i] = views.AvlRecords{
			Plate:          result.Plate,
			Imei:           result.Imei,
			Latitude:       result.Latitude,
			Longitude:      result.Longitude,
			TimeStampEvent: result.Timestamp,
			Event:          result.Event,
			Speed:          result.Speed,
			Id_company:     &Id_company,
			Company:        &Company,
			Properties:     string(propertiesJson),
		}
	}
	// Obtener las relaciones de los clientes
	vehiclesToCustomer, err := GetCustomerRelations(db.DBConn, avlRecords)
	if err != nil {
		return nil, err
	}

	// Completar los campos Fk_customer y Customer con el resultado de GetCustomerRelations
	for i, record := range avlRecords {
		for _, vtc := range vehiclesToCustomer {
			if *vtc.Plate == *record.Plate {
				avlRecords[i].Id_customer = vtc.FkCustomer
				avlRecords[i].Customer = vtc.Customer
				break
			}
		}
	}

	return avlRecords, nil
}
