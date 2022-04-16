package controller

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	temperature, _ := t.TemperatureModel.GetTemperatureByMonth(payload.Month, payload.Year)
	csvFile, err := processCsv(temperature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var filename string
	if filename, err = processImage(payload.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("Filename : ", filename)
	zipFile, _ := processZipFile([]string{csvFile, filename})
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", "attachment; filename='data.zip'")
	c.File(zipFile)
}

func processCsv(temperature []*entity.Temperature) (string, error) {
	fileName := fmt.Sprintf("data/temperature-data-%v-%v.csv", time.Now().Month(), time.Now().Year())
	csvFile, err := os.Create(fileName)
	defer csvFile.Close()
	if err != nil {
		return "", err
	}
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()
	var data [][]string
	for _, record := range temperature {
		temperatureStr := fmt.Sprintf("%f", record.Temperature)
		row := []string{strconv.Itoa(int(record.ID)), temperatureStr, record.TimeCreated.String()}
		data = append(data, row)
	}
	err = csvWriter.WriteAll(data)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func processZipFile(filepaths []string) (string, error) {
	log.Println("creating zip archive...")
	zipName := fmt.Sprintf("download/archive-%v-%v.zip", time.Now().Month(), time.Now().Year())
	archive, err := os.Create(zipName)
	if err != nil {
		log.Panicln(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	log.Println("opening first file...")
	f1, err := os.Open(filepaths[0])
	if err != nil {
		log.Panicln(err)
	}
	defer f1.Close()

	fmt.Println("writing first file to archive...")
	w1, err := zipWriter.Create("download/" + filepaths[0])
	if err != nil {
		log.Panicln(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		log.Panicln(err)
	}

	log.Println("opening second file")
	f2, err := os.Open(filepaths[1])
	if err != nil {
		log.Panicln(err)
	}
	defer f2.Close()

	fmt.Println("writing second file to archive...")
	w2, err := zipWriter.Create("download/" + filepaths[1])
	if err != nil {
		log.Panicln(err)
	}
	if _, err := io.Copy(w2, f2); err != nil {
		log.Panicln(err)
	}
	log.Println("closing zip archive...")
	zipWriter.Close()
	return zipName, nil
}

func processImage(base string) (string, error) {
	index := strings.Index(base, ";base64,")
	if index < 0 {
		log.Panicln("Invalid Image")
		return "", errors.New("Invalid file")
	}
	imageType := base[11:index]
	log.Println("Image type : ", imageType)
	unbased, _ := base64.StdEncoding.DecodeString(base[index+8:])
	r := bytes.NewReader(unbased)
	switch imageType {
	case "png":
		filename := fmt.Sprintf("data/Image-%v-%v.png", time.Now().Month(), time.Now().Year())
		img, err := png.Decode(r)
		if err != nil {
			log.Panicln("Bad PNG, ", err.Error())
			return "", err
		}
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Panicln("Cannot open file")
			return "", err
		}
		png.Encode(file, img)
		return filename, nil
	case "jpg":
		filename := fmt.Sprintf("data/Image-%v-%v.png", time.Now().Month(), time.Now().Year())
		img, err := jpeg.Decode(r)
		if err != nil {
			log.Panicln("Bad JPG, ", err.Error())
			return "", err
		}
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Panicln("Cannot open file")
			return "", err
		}
		jpeg.Encode(file, img, nil)
		return filename, nil
	}
	return "", errors.New("cant infer image type")
}
