package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
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
	TemperatureModel         model.ITemperatureModel
	Client                   *mqtt.Client
	db                       *gorm.DB
	telegramController       ITelegramController
	lastTimeOutsideSetPoint  time.Time
	isOutOfRangeConditionMet bool
	isEvacuationConditionMet bool
}

func CreateMqttController(model model.ITemperatureModel, client *mqtt.Client, db *gorm.DB, telegramController ITelegramController) IMqttController {
	return &MqttController{
		TemperatureModel:         model,
		Client:                   client,
		db:                       db,
		telegramController:       telegramController,
		isOutOfRangeConditionMet: false,
		isEvacuationConditionMet: false,
	}
}

func (mc *MqttController) TemperatureProcessor(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message received from %v\n", message.Topic())
	var temperature entity.Temperature
	payload := message.Payload()
	if err := json.Unmarshal(payload, &temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	temperature.TimeCreated = time.Now()
	if err := mc.TemperatureModel.StoreTemperature(&temperature); err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
	parsedTemperature, _ := strconv.ParseFloat(temperature.Temperature, 32)
	if mc.isEvacuationConditionMet {
		fmt.Println("Evacuation Routine")
		evacuationCode := mc.telegramController.Evacuation()
		if evacuationCode >= 4 {
			mc.isEvacuationConditionMet = false
		}
	} else if parsedTemperature < 2 || parsedTemperature > 8 && mc.isOutOfRangeConditionMet {
		fmt.Println("Out of range routine")
		mc.telegramController.OutOfRangeTemperature(temperature)
		mc.isOutOfRangeConditionMet = false
		mc.isEvacuationConditionMet = true
	} else if parsedTemperature < 3.3 || parsedTemperature > 4.3 {
		mc.temperatureAnomaly(temperature)
	}
}

func (mc *MqttController) temperatureAnomaly(temperature entity.Temperature) {
	if (mc.lastTimeOutsideSetPoint == time.Time{}) {
		mc.lastTimeOutsideSetPoint = time.Now()
		fmt.Println("Door Opened Routine")
		mc.telegramController.DoorOpened()
	} else if time.Now().Sub(mc.lastTimeOutsideSetPoint) < 12*time.Second {
		fmt.Printf("Time : %v\n", mc.lastTimeOutsideSetPoint)
		fmt.Printf("Time elapsed : %v\n", time.Now().Sub(mc.lastTimeOutsideSetPoint))
		fmt.Println("Door Opened Routine | Below 15 minute")
		mc.telegramController.DoorOpened()
	} else if time.Now().Sub(mc.lastTimeOutsideSetPoint) >= 12*time.Second &&
		time.Now().Sub(mc.lastTimeOutsideSetPoint) < 24*time.Second {
		fmt.Println("Early Warning Routine")
		mc.telegramController.EarlyWarningProcessor(temperature)
	} else {
		fmt.Println("Out of range condition is true")
		mc.isOutOfRangeConditionMet = true
	}
}
