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
	userModel := model.CreateUserModel(db)
	temperatureModel := model.CreateTemperatureModel(db)
	userController := controller.CreateUserController(userModel)
	temperatureController := controller.CreateTemperatureController(temperatureModel)
	telegramController := controller.CreateTelegramController(telegramClient)
	mqttController := controller.CreateMqttController(temperatureModel, &mqttClient, db, telegramController)

	router.GET(ROOT+"user", userController.GetUsers)
	router.GET(ROOT+"user/:id", userController.GetUserById)
	router.POST(ROOT+"user", userController.CreateUser)
	router.PATCH(ROOT+"user/:id", userController.UpdateUser)
	router.DELETE(ROOT+"user/:id", userController.DeleteUser)

	router.POST(ROOT+"temperature", temperatureController.GetTemperatureByMonth)
	router.POST(ROOT+"temperature/download", temperatureController.GetTemperatureFile)

	mqttClient.Subscribe("ahaha_JsonData", QOS, mqttController.TemperatureProcessor)

	err := router.Run(":8080")
	if err != nil {
		panic("Error. Reason : " + err.Error())
	}
}
