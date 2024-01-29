package service

import (
	"SimonBK_ControlSat/api/views"
	"SimonBK_ControlSat/domain/models"
	"SimonBK_ControlSat/infra/db"
	"fmt"
	"strings"
)

func GetForCustomer(FkCompany *int, imei []string) ([]views.ResultSqlServer, error) {
	results := []views.ResultSqlServer{}

	// Obtener la consulta correspondiente a FkCompany
	query, ok := models.Consultas[*FkCompany] // Usar el mapa Consultas del paquete models
	if !ok {
		return nil, fmt.Errorf("FkCompany inválido: %v", *FkCompany)
	}

	// Convertir la lista de IMEIs en una cadena para usar en la consulta SQL
	imeiList := "'" + strings.Join(imei, "','") + "'"

	// Reemplazar un marcador de posición en la consulta con la lista de IMEIs
	query = strings.Replace(query, "{IMEI}", imeiList, -1)

	rows, err := db.SQLServerConn.Raw(query).Rows() // Raw SQL
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r views.ResultSqlServer
		err = rows.Scan(&r.Imei, &r.Plate, &r.Latitude, &r.Longitude, &r.Timestamp, &r.Event)
		if err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	return results, nil
}
