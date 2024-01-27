package service

import (
	"SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/domain/models"
	"database/sql"
	"log"
)

func GetAll(db *sql.DB) ([]views.ResultSqlServer, error) {
	var results []views.ResultSqlServer

	// Define el orden en el que quieres ejecutar las consultas
	order := []int{12, 2, 4, 11}

	for _, index := range order {
		// Asegúrate de que el índice esté dentro del rango de models.Consultas
		if index < 0 || index >= len(models.ConsultasAll) {
			log.Printf("Index out of range: %v", index)
			continue
		}

		query := models.ConsultasAll[index]
		rows, err := db.Query(query)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			continue
		}
		defer rows.Close()

		for rows.Next() {
			var result views.ResultSqlServer
			err := rows.Scan(&result.Imei, &result.Plate, &result.Description, &result.Latitude, &result.Longitude, &result.Timestamp, &result.Event)
			if err != nil {
				log.Printf("Error scanning row: %v", err)
				continue
			}

			results = append(results, result)
		}

		if err := rows.Err(); err != nil {
			log.Printf("Error with rows: %v", err)
		}
	}

	return results, nil
}
