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

	for _, key := range order {
		// Asegúrate de que la clave exista en models.ConsultasAll
		query, ok := models.ConsultasAll[key]
		if !ok {
			log.Printf("Key not found: %v", key)
			continue
		}

		rows, err := db.Query(query)
		if err != nil {
			log.Printf("Error executing query: %v", err)
			continue
		}
		defer rows.Close()

		for rows.Next() {
			var result views.ResultSqlServer
			err := rows.Scan(&result.Imei, &result.Plate, &result.Description, &result.Speed, &result.Latitude, &result.Longitude, &result.Timestamp, &result.Event, &result.Odometer)
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
