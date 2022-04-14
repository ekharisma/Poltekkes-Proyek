package controller

import (
	"encoding/base64"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"github.com/gin-gonic/gin"
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
	var payload entity.DownloadRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = t.TemperatureModel.GetTemperatureByMonth(payload.Month, payload.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err = processImage(payload.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusAccepted)
}

func processImage(image string) error {
	splitBase := strings.Split(image, ",")
	if strings.Contains(splitBase[0], "png") {
		// err := convertBaseToPng(splitBase[1])
		return nil
	} else if strings.Contains(splitBase[0], "jpg") {
		err := convertBaseToJpg(splitBase[1])
		return err
	} else {
		return errors.New("file not supported")
	}
}

func convertBaseToJpg(s string) error {
	jpgDecoded := base64.NewDecoder(base64.StdEncoding.WithPadding(base64.NoPadding), strings.NewReader(strings.TrimSpace(s)))
	jpgImage, err := jpeg.Decode(jpgDecoded)
	if err != nil {
		fmt.Println("Error decoded jpg. Reason : ", err.Error())
		return err
	}
	file, err := os.Create("images/image.jpg")
	if err != nil {
		fmt.Println("Error creating jpg. Reason : ", err.Error())
		return err
	}
	err = png.Encode(file, jpgImage)
	if err != nil {
		fmt.Println("Error encoding jpg. Reason : ", err.Error())
		return err
	}
	defer file.Close()
	return nil
}

// func convertBaseToPng(s string) error {
// 	pngDecoded, _ := base64.StdEncoding.DecodeString(s)
// 	pngImage, err := png.Decode(pngDecoded)
// 	if err != nil {
// 		fmt.Println("Error decoded png. Reason : ", err.Error())
// 		return err
// 	}
// 	file, err := os.Create("Coba.png")
// 	if err != nil {
// 		fmt.Println("Error creating png. Reason : ", err.Error())
// 		return err
// 	}
// 	err = png.Encode(file, pngImage)
// 	if err != nil {
// 		fmt.Println("Error encoding png.  Reason : ", err.Error())
// 		return err
// 	}
// 	defer file.Close()
// 	return nil
// }
