package controllers

import (
	"SimonBK_ControlSat/domain/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary Obtiene ControlSat
// @Description Obtine ControlSat
// @Tags Controlsat
// @Produce json
// @Success 200 {object} swagger.GormModelStub
// @Failure 500 {object} swagger.ErrorResponse
// @Router /Finandina [get]
func GetAllFinandinaController(c *gin.Context) {
	results, err := service.GetAllFinandina()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
