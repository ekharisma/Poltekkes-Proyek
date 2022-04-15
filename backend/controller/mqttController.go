package controller

import (
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"gorm.io/gorm"
)

type IMqttController interface {
	TemperatureProcessor(client mqtt.Client, message mqtt.Message)
}

type MqttController struct {
	TemperatureModel   model.ITemperatureModel
	Client             *mqtt.Client
	db                 *gorm.DB
	telegramController ITelegramController
}

func CreateMqttController(model model.ITemperatureModel, client *mqtt.Client, db *gorm.DB, telegramController ITelegramController) IMqttController {
	return &MqttController{
		TemperatureModel:   model,
		Client:             client,
		db:                 db,
		telegramController: telegramController,
	}
}

func (mc *MqttController) TemperatureProcessor(client mqtt.Client, message mqtt.Message) {
	var temperature entity.Temperature
	payload := message.Payload()
	if err := json.Unmarshal(payload, &temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	temperature.TimeCreated = time.Now()
	if err := mc.TemperatureModel.StoreTemperature(&temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	if temperature.Temperature < 2.0 || temperature.Temperature > 8.0 {
		mc.telegramController.ProcessMessageRequest(temperature)
	}
}
