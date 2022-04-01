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
	service.CreatePostgresClient(constant.DBHost, constant.DBUsername, constant.DBPassword, constant.DBPort, constant.DBName)
	service.CreateMqttClient(constant.Broker, constant.MqttPort)
	router := gin.Default()
	db := service.GetPostgresDBClient()
	//mqttClient := service.GetMqttClient()
	userModel := model.CreateUserModel(db)
	userController := controller.CreateUserController(userModel)

	router.GET(ROOT+"user", userController.GetUsers)
	router.GET(ROOT+"user/:id", userController.GetUserById)
	router.POST(ROOT+"user", userController.CreateUser)
	router.PATCH(ROOT+"user/:id", userController.PatchUser)
	router.DELETE(ROOT+"user/:id", userController.DeleteUser)
	router.Run(":8080")
}
