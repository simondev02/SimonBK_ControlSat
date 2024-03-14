package service

import (
	views "SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/domain/models"
	"SimonBK_ControlSat/infra/db"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetAllFinandina(FkCompany int) ([]views.ResultSqlServer, error) {

	results := []views.ResultSqlServer{}

	// Obtener la consulta correspondiente a FkCompany
	query, ok := models.ConsultasAll[FkCompany] // Usar el mapa Consultas del paquete models
	if !ok {
		return nil, fmt.Errorf("FkCompany inválido: %v", FkCompany)
	}

	db.SQLServerConn = db.SQLServerConn.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})

	rows, err := db.SQLServerConn.Raw(query).Rows() // Raw SQL
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r views.ResultSqlServer
		err = rows.Scan(&r.Imei, &r.Plate, &r.Description, &r.Speed, &r.Latitude, &r.Longitude, &r.Timestamp, &r.Event, &r.Odometer)
		if err != nil {
			return nil, err
		}
		// Agregar 5 horas a Timestamp
		temp := r.Timestamp.Add(5 * time.Hour)
		r.Timestamp = &temp

		// Cambiar el valor de Event si es "acc"
		if *r.Event == "acc " {
			*r.Event = "Por tiempo"
		}

		results = append(results, r)
	}

	return results, nil
}
