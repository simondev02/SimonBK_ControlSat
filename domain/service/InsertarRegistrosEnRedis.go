package service

import (
	"SimonBK_ControlSat/api/views"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func InsertarRegistrosEnRedis(client *redis.Client, records []views.AvlRecords) error {
	for _, record := range records {
		err := InsertarEnRedis(client, record)
		if err != nil {
			return fmt.Errorf("error al insertar el registro en Redis: %v", err)
		}
	}
	return nil
}
