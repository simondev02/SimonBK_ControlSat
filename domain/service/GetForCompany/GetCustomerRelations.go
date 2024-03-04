package service

import (
	"SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/domain/models"

	"gorm.io/gorm"
)

type VehiclesToCustomer struct {
	Plate      *string
	FkCustomer *int
	Customer   *string
}

func GetCustomerRelations(db *gorm.DB, records []views.AvlRecords) ([]VehiclesToCustomer, error) {
	var vehicles []models.Vehicle

	// Extraer las placas de los registros
	plates := make([]string, len(records))
	for i, record := range records {
		if record.Plate != nil {
			plates[i] = *record.Plate
		}
	}

	// Buscar los vehículos que tienen esas placas y precargar la relación Customer
	err := db.Preload("Customer").Where("plate IN (?)", plates).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	// Crear un slice de VehiclesToCustomer
	vehiclesToCustomer := make([]VehiclesToCustomer, len(vehicles))
	for i, vehicle := range vehicles {
		vehiclesToCustomer[i] = VehiclesToCustomer{
			Plate:      vehicle.Plate,
			FkCustomer: vehicle.FkCustomer,
			Customer:   vehicle.Customer.Name,
		}
	}

	return vehiclesToCustomer, nil
}
