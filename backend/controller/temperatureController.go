package controller

import (
	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ITemperatureController interface {
	GetTemperatureByMonth(c *gin.Context)
	GetTemperatureFile(c *gin.Context)
}

type TemperatureController struct {
	TemperatureModel model.ITemperatureModel
}

func CreateTemperatureController(model model.ITemperatureModel) ITemperatureController {
	return &TemperatureController{
		TemperatureModel: model,
	}
}

func (t *TemperatureController) GetTemperatureByMonth(c *gin.Context) {
	var payload entity.DateRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	temperatureByMonth, err := t.TemperatureModel.GetTemperatureByMonth(payload.Month, payload.Year)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"payload": temperatureByMonth,
	})
}

func (t *TemperatureController) GetTemperatureFile(c *gin.Context) {

}
