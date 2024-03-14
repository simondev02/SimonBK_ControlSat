package service

import (
	views "SimonBK_ControlSat/api/views"
	GetAllForCompany "SimonBK_ControlSat/domain/service/GetForCompany"
	"SimonBK_ControlSat/infra/db"
	"fmt"
	"log"
	"sync"
)

func GetAllInReddis() error {
	redisClient, err := db.CreateRedisClient()
	if err != nil {
		return fmt.Errorf("error al conectar a Redis: %w", err)
	}

	var wg sync.WaitGroup
	companies := []func() ([]views.AvlRecords, error){
		GetAllForCompany.GetAllCarfiao,
		GetAllForCompany.GetAllFinnan,
		GetAllForCompany.GetAllFZ,
		GetAllForCompany.GetAllPresAuto,
	}

	for _, company := range companies {
		wg.Add(1)
		go func(company func() ([]views.AvlRecords, error)) {
			defer wg.Done()
			records, err := company()
			if err != nil {
				log.Printf("error al obtener registros: %v", err)
				return
			}

			err = InsertarRegistrosEnRedis(redisClient, records)
			if err != nil {
				log.Printf("error al insertar registros en Redis: %v", err)
			}
		}(company)
	}

	wg.Wait()
	return nil
}
