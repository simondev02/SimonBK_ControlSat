package service

import (
	views "SimonBK_ControlSat/api/views"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func InsertarRegistrosEnRedis(client *redis.Client, records []views.AvlRecords) error {

	if len(records) == 0 {
		log.Println("[ControlSat] - No hay registros para insertar en Redis")
		return nil
	}
	var count int
	for _, record := range records {
		err := InsertarEnRedis(client, record)
		if err != nil {
			return fmt.Errorf("error al insertar el registro en Redis: %v", err)
		}
		count++
	}
	if len(records) > 0 {
		log.Println("[ControlSat] - Inserciones en Redis", count)
	}
	return nil
}
