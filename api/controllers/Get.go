package controllers

import (
	"SimonBK_ControlSat/domain/service"
	"SimonBK_ControlSat/infra/db"
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
// @Router /ControlSat/Results [get]
func Get(c *gin.Context) {

	fkCompanyvalue, exists := c.Get("FkCompany")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FkCompany no proporcionado"})
		return
	}

	fkCustomervalue, exists := c.Get("FkCustomer")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "FkCustomer no proporcionado"})
		return
	}

	fkCompany := fkCompanyvalue.(int)
	fkCustomer := fkCustomervalue.(int)

	// Obtén la conexión a la base de datos SQL Server
	sqlDB, err := db.SQLServerConn.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imeis, err := service.Get(db.DBConn, sqlDB, &fkCompany, &fkCustomer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, imeis)

}
