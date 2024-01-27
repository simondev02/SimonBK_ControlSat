package controllers

import (
	"SimonBK_ControlSat/domain/service"
	"SimonBK_ControlSat/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router /getAll [get]
func GetAll(c *gin.Context) {
	// Obtén la conexión a la base de datos SQL Server
	sqlDB, err := db.SQLServerConn.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results, err := service.GetAll(sqlDB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, results)
}
