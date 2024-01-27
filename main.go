// @SecurityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"SimonBK_ControlSat/docs"
	"SimonBK_ControlSat/infra/db"
	"SimonBK_ControlSat/routers"

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

	// // Establecer la conexión con Redis
	// redisClient, err := db.CreateRedisClient()
	// if err != nil {
	// 	fmt.Println("Error al conectar con Redis:", err)
	// 	return
	// } else {
	// 	fmt.Println("Conexión exitosa en Redis")
	// }

	// //Iniciar la goroutine para almacenar los resultados en Redis
	// go func() {
	// 	FkCompany := 12
	// 	FkCustomer := 13337
	// 	results, err := service.GetResultsWithNewStruct(&FkCompany, &FkCustomer)
	// 	if err != nil {
	// 		fmt.Println("Error al obtener los resultados:", err)
	// 		return
	// 	}

	// 	err = service.StoreNewStructInRedis(redisClient, results)
	// 	if err != nil {
	// 		fmt.Println("Error al almacenar los resultados en Redis:", err)
	// 		return
	// 	}
	// }()

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
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST" + ":" + os.Getenv("SERVICE_PORT"))
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
	err = r.Run(":" + os.Getenv("SERVICE_PORT")) // escucha y sirve en 0.0.0.0:60031  (por defecto)

	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
}
