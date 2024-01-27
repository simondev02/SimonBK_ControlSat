package service

import (
	"SimonBK_ControlSat/domain/models"

	"gorm.io/gorm"
)

func GetPostgresImei(db *gorm.DB, FkCompany *int, FkCustomer *int) ([]string, error) {
	var vehicles []models.Vehicle
	var results []string

	query := db.
		Preload("AvlDevice").
		Preload("Company").
		Preload("Customer").
		Limit(2000)

	switch {

	// Ahora, groupedVehicles es un mapa donde la clave es el FkCompany y el valor es una lista de vehículos
	case *FkCompany != 0 && *FkCustomer == 0:
		query = query.Where("fk_company = ?", FkCompany)
	case *FkCompany != 0 && *FkCustomer != 0:
		query = query.Where("fk_customer = ?", FkCustomer)
	}

	err := query.Find(&vehicles).Error
	if err != nil {
		return nil, err
	}

	for _, vehicle := range vehicles {
		if vehicle.AvlDevice.IMEI != nil {
			results = append(results, *vehicle.AvlDevice.IMEI)
		} else {
			// Manejar el caso en que vehicle.AvlDevice.IMEI es nil
			// Puedes registrar un error o devolver un error desde la función
		}
	}

	return results, nil
}
