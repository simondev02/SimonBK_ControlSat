package routers

import (
	"SimonBK_ControlSat/api/controllers"
	"SimonBK_ControlSat/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {

	// Validacion de acces Token
	r.Use(middleware.ValidateTokenMiddleware())

	// Grupo de rutas para vehiculos
	ControSat := r.Group("/ControlSat")
	{
		ControSat.GET("/Results/", controllers.Get)

		ControSat.GET("GetAll/", controllers.GetAll)
		ControSat.GET("/finandina", controllers.GetAllFinandinaController)

	}

}
