package controller

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"gorm.io/gorm"
	"log"
	"time"
)

type IMqttController interface {
	TemperatureProcessor(client mqtt.Client, message mqtt.Message)
}

type MqttController struct {
	TemperatureModel model.ITemperatureModel
	Client           *mqtt.Client
	db               *gorm.DB
	emailController  IEmailController
}

func CreateMqttController(model model.ITemperatureModel, client *mqtt.Client, db *gorm.DB, emailController IEmailController) IMqttController {
	return &MqttController{
		TemperatureModel: model,
		Client:           client,
		db:               db,
		emailController:  emailController,
	}
}

func (mc *MqttController) TemperatureProcessor(client mqtt.Client, message mqtt.Message) {
	fmt.Println("Get Message from :", message.Topic())
	var temperature entity.Temperature
	payload := message.Payload()
	if err := json.Unmarshal(payload, &temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	temperature.TimeCreated = time.Now()
	fmt.Println("Temperature", temperature)
	if err := mc.TemperatureModel.StoreTemperature(&temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	if isTemperatureAnomaly(temperature.Temperature) {
		fmt.Println("Temperature Anomaly Detected!")
		go mc.emailController.SendEmail(temperature)
	}
}

func isTemperatureAnomaly(temperature float64) bool {
	if temperature < 2.0 {
		return true
	} else if temperature > 8.0 {
		return true
	}
	return false
}
