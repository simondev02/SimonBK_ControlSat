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
		// Agregar 5 horas a Timestamp y formatear
		temp, err := time.Parse("2006-01-02T15:04:05Z", r.Timestamp.Add(5*time.Hour).Format("2006-01-02T15:04:05Z"))
		if err != nil {
			return nil, err
		}
		r.Timestamp = temp

		results = append(results, r)
	}

	return results, nil
}
