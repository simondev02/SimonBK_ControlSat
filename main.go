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
	"SimonBK_ControlSat/docs"
	"SimonBK_ControlSat/infra/db"
	"SimonBK_ControlSat/routers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Establecer la conexión con la base de datos PostgreSQL
	err := db.ConnectDB()

	if err != nil {
		fmt.Println("Error al conectar con la base de datos PostgreSQL:", err)
		return
	} else {
		fmt.Println("Conexión exitosa en PostgreSQL")
	}

	// Establecer la conexión con la base de datos SQL Server
	err = db.ConnectSQLServer()

	if err != nil {
		fmt.Println("Error al conectar con la base de datos SQL Server:", err)
		return
	} else {
		fmt.Println("Conexión exitosa en SQL Server")
	}

	// Configurar CORS
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Configurar Swagger
	docs.SwaggerInfo.Title = "API vehicle"
	docs.SwaggerInfo.Description = "Probar y confirmar el funcionamiento correcto del microServicio vehicle, la creación, visualización, modificación y eliminación de un registro."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = "/"

	// Configurar e iniciar el enrutador
	routers.SetupRouter(r)

	// Agregar la ruta de Swagger sin el middleware de validación de token
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configurar la señal de captura
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		// Código de limpieza: cierra la conexión a la base de datos
		db.CloseDB()
		db.CloseSQLServerDB()
		os.Exit(0)
	}()

	// Escuchar y servir

	certFile := os.Getenv("TLS_CERT")
	certKey := os.Getenv("TLS_KEY")
	if certFile == "" || certKey == "" {
		log.Println("Error al leer las variables de entorno.")
		db.CloseDB()
		os.Exit(1)
	}

	err = r.RunTLS(":"+os.Getenv("SERVICE_PORT"), certFile, certKey) // escucha y sirve en 0.0.0.0:60031  (por defecto)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
}
