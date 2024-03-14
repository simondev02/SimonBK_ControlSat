// @API ControlSat
// @version 1
// @BasePath /ControlSat
// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @SecurityDefinitions.apikey ApiKeyAuth2
// @in header
// @name X-API-KEY

package main

import (
	"SimonBK_ControlSat/domain/service"
	"SimonBK_ControlSat/infra/db"

	"log"
	"time"

	"os"
	"os/signal"
	"syscall"
)

func main() {

	db.ConnectDB()
	db.ConnectSQLServer()
	service.GetAllInReddis()

	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			err := service.GetAllInReddis()
			if err != nil {
				log.Printf("[ControlSat] - Error al ejecutar el servicio: %v", err)
			}
		}
	}()
	defer ticker.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		db.CloseDB()
		db.CloseSQLServerDB()
		os.Exit(0)
	}()

	for {
	}
}
