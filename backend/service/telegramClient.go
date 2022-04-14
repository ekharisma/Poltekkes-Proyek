package service

import (
	"github.com/ekharisma/poltekkes-webservice/constant"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

var telegramClient *tgbotapi.BotAPI
var telegramClientSingleton sync.Once
var err error

func CreateTelegramBotClient() {
	telegramClientSingleton.Do(func() {
		telegramClient, err = tgbotapi.NewBotAPI(constant.TelegramToken)
		log.Println("Authorized on account : ", telegramClient.Self.UserName)
		if err != nil {
			log.Panicln("Error. Reason : ", err.Error())
		}
	})
}

func GetTelegramBotClient() *tgbotapi.BotAPI {
	return telegramClient
}
