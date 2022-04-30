package main

import (
	"github.com/ekharisma/poltekkes-webservice/constant"
	"github.com/ekharisma/poltekkes-webservice/controller"
	"github.com/ekharisma/poltekkes-webservice/model"
	"github.com/ekharisma/poltekkes-webservice/service"
	"github.com/gin-gonic/gin"
)

func main() {
	const ROOT = "/api/"
	const QOS = 2
	service.CreateTelegramBotClient()
	service.CreatePostgresClient(constant.DBHost, constant.DBUsername, constant.DBPassword, constant.DBPort, constant.DBName)
	service.CreateMqttClient(constant.Broker, constant.MqttPort)
	router := gin.Default()
	db := service.GetPostgresDBClient()
	mqttClient := service.GetMqttClient()
	telegramClient := service.GetTelegramBotClient()
	temperatureModel := model.CreateTemperatureModel(db)
	temperatureController := controller.CreateTemperatureController(temperatureModel)
	telegramController := controller.CreateTelegramController(telegramClient)
	mqttController := controller.CreateMqttController(temperatureModel, &mqttClient, db, telegramController)

	router.POST(ROOT+"temperature", temperatureController.GetTemperatureByMonth)
	router.POST(ROOT+"temperature/download", temperatureController.GetTemperatureFile)

	mqttClient.Subscribe("ahaha/#", QOS, mqttController.TemperatureProcessor)

	err := router.Run(":8080")
	if err != nil {
		panic("Error. Reason : " + err.Error())
	}
}
