package db

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var SQLServerConn *gorm.DB

func ConnectSQLServer() error {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error al leer variables de entorno", err)
	}

	dbHost := os.Getenv("SQL_SERVER_HOST")
	dbPort := os.Getenv("SQL_SERVER_PORT")
	dbUser := os.Getenv("SQL_SERVER_USER")
	dbPass := os.Getenv("SQL_SERVER_PASSWORD")
	dbName := os.Getenv("SQL_SERVER_DB_NAME")
	password := url.QueryEscape(dbPass)
	DSN := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, password, dbHost, dbPort, dbName)

	SQLServerConn, err = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("SQL Server DB CONNECTED")
	}
	return err
}

func CloseSQLServerDB() error {
	sqlDB, err := SQLServerConn.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	fmt.Println("SQL Server DB DISCONNECTED")

	return nil
}
