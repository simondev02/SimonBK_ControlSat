package controllers

import (
	"SimonBK_ControlSat/domain/service"
	"SimonBK_ControlSat/infra/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth2
// @Summary Obtiene ControlSat
// @Description Obtine ControlSat
// @Tags Finandina
// @Produce json
// @Success 200 {object} swagger.GormModelStub
// @Failure 500 {object} swagger.ErrorResponse
// @Router /Finandina [get]
func GetAllFinandinaController(c *gin.Context) {
	results, err := service.GetAllRecords(db.DBConn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
