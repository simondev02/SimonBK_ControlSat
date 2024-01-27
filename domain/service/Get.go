package service

import (
	"SimonBK_ControlSat/api/views"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

func Get(db *gorm.DB, sqlDB *sql.DB, FkCompany *int, FkCustomer *int) ([]views.ResultSqlServer, error) {
	// Si FkCompany y FkCustomer son 0, llamar a la función GetAll
	if *FkCompany == 0 && *FkCustomer == 0 {
		results, err := GetAll(sqlDB)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	// Llamar a la función GetPostgresImei con los parámetros FkCompany y FkCustomer
	imeis, err := GetPostgresImei(db, FkCompany, FkCustomer)
	if err != nil {
		return nil, err
	}

	// Si FkCompany y FkCustomer no son nil, llamar a la función GetForCustomer
	if FkCompany != nil && FkCustomer != nil {
		results, err := GetForCustomer(FkCompany, imeis)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	// Si FkCompany no es nil y FkCustomer es nil, llamar a la función GetForCompany
	if FkCompany != nil && FkCustomer == nil {
		results, err := GetForCompany(FkCompany, imeis)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	// Si FkCompany y FkCustomer son nil, devolver un error
	return nil, errors.New("FkCompany y FkCustomer no pueden ser nil")
}
