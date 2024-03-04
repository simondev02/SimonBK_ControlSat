package controllers

import (
	"SimonBK_ControlSat/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	// Llamar a la función GetAllInReddis
	allRecords := service.GetAllInReddis()

	c.JSON(http.StatusOK, allRecords)
}
