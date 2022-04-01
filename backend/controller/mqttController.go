package controller

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"gorm.io/gorm"
	"log"
)

type IMqttController interface {
	TemperatureProcessor(client mqtt.Client, message mqtt.Message)
}

type MqttController struct {
	TemperatureModel model.ITemperatureModel
	Client           *mqtt.Client
	db               *gorm.DB
}

func CreateMqttController(model model.ITemperatureModel, client *mqtt.Client, db *gorm.DB) IMqttController {
	return &MqttController{
		TemperatureModel: model,
		Client:           client,
		db:               db,
	}
}

func (mc *MqttController) TemperatureProcessor(client mqtt.Client, message mqtt.Message) {
	var temperature entity.Temperature
	payload := message.Payload()
	if err := json.Unmarshal(payload, &temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	if err := mc.TemperatureModel.StoreTemperature(&temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	//if isTemperatureAnomaly(temperature.Temperature) {go mc.}
}

func isTemperatureAnomaly(temperature float64) bool {
	if temperature < 2.0 {
		return true
	} else if temperature > 8.0 {
		return true
	}
	return false
}
