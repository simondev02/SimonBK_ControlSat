package service

import (
	"SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/domain/models"
	"SimonBK_ControlSat/infra/db"
	"fmt"
	"time"
)

func GetAllFinandina() ([]views.ResultSqlServer, error) {

	var FkCompany = 12

	results := []views.ResultSqlServer{}

	// Obtener la consulta correspondiente a FkCompany
	query, ok := models.ConsultasAll[FkCompany] // Usar el mapa Consultas del paquete models
	if !ok {
		return nil, fmt.Errorf("FkCompany inválido: %v", FkCompany)
	}

	rows, err := db.SQLServerConn.Raw(query).Rows() // Raw SQL
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r views.ResultSqlServer
		err = rows.Scan(&r.Imei, &r.Plate, &r.Description, &r.Speed, &r.Latitude, &r.Longitude, &r.Timestamp, &r.Event)
		if err != nil {
			return nil, err
		}
		// Agregar 5 horas a Timestamp
		temp := r.Timestamp.Add(5 * time.Hour)
		r.Timestamp = &temp

		results = append(results, r)
	}

	return results, nil
}
