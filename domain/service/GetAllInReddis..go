package service

import (
	GetAllForCompany "SimonBK_ControlSat/domain/service/GetForCompany"
	"SimonBK_ControlSat/infra/db"
	"fmt"
)

func GetAllInReddis() error {
	// Insertamos la trama en Redis
	redisClient, err := db.CreateRedisClient()
	if err != nil {
		return fmt.Errorf("Error al conectar a Redis: %w", err)
	}

	// Primera función
	recordsCarfiao, err := GetAllForCompany.GetAllCarfiao()
	if err != nil {
		return fmt.Errorf("Error al obtener registros de Carfiao: %w", err)
	}
	err = InsertarRegistrosEnRedis(redisClient, recordsCarfiao)
	if err != nil {
		return fmt.Errorf("Error al insertar registros de Carfiao en Redis: %w", err)
	}

	// Segunda función
	recordsFinnan, err := GetAllForCompany.GetAllFinnan()
	if err != nil {
		return fmt.Errorf("Error al obtener registros de Finnan: %w", err)
	}
	err = InsertarRegistrosEnRedis(redisClient, recordsFinnan)
	if err != nil {
		return fmt.Errorf("Error al insertar registros de Finnan en Redis: %w", err)
	}

	// Tercera función
	recordsFZ, err := GetAllForCompany.GetAllFZ()
	if err != nil {
		return fmt.Errorf("Error al obtener registros de FZ: %w", err)
	}
	err = InsertarRegistrosEnRedis(redisClient, recordsFZ)
	if err != nil {
		return fmt.Errorf("Error al insertar registros de FZ en Redis: %w", err)
	}

	// Cuarta función
	recordsPresAuto, err := GetAllForCompany.GetAllPresAuto()
	if err != nil {
		return fmt.Errorf("Error al obtener registros de PresAuto: %w", err)
	}
	err = InsertarRegistrosEnRedis(redisClient, recordsPresAuto)
	if err != nil {
		return fmt.Errorf("Error al insertar registros de PresAuto en Redis: %w", err)
	}

	return nil
}
