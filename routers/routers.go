package routers

import (
	"SimonBK_ControlSat/api/controllers"
	"SimonBK_ControlSat/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {

	// Grupo de rutas para Finandina
	finandina := r.Group("/Finandina")
	{
		finandina.Use(middleware.ValidateAPIKey()) // Añadir el middleware de validación de API key aquí
		finandina.GET("/", controllers.GetAllFinandinaController)
	}

	// Grupo de rutas que requieren validación de token
	authorized := r.Group("/")
	authorized.Use(middleware.ValidateTokenMiddleware())
	{
		// Grupo de rutas para vehiculos
		ControSat := authorized.Group("/ControlSat")
		{
			ControSat.GET("/Results/", controllers.Get)
		}
	}
}
