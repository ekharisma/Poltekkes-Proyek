package controller

import (
	"fmt"
	"github.com/ekharisma/poltekkes-webservice/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TelegramController struct {
	telegramClient *tgbotapi.BotAPI
}

type ITelegramController interface {
	SendMessage(temperature entity.Temperature)
}

func CreateTelegramController(telegramClient *tgbotapi.BotAPI) ITelegramController {
	return &TelegramController{
		telegramClient: telegramClient,
	}
}

func constructMessage(temperature entity.Temperature) string {
	return fmt.Sprintf(`Pada %v, suhu diluar batas normal dengan %v`, temperature.Timestamp, temperature.Temperature)
}

func (c TelegramController) SendMessage(temperature entity.Temperature) {
	c.telegramClient.Debug = true
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60
	updatesChannel := c.telegramClient.GetUpdatesChan(update)
	for updates := range updatesChannel {
		if updates.Message != nil {
			message := tgbotapi.NewMessage(updates.Message.Chat.ID, constructMessage(temperature))
			message.ReplyToMessageID = updates.Message.MessageID
			_, err := c.telegramClient.Send(message)
			if err != nil {
				log.Panicln("Error", err.Error())
			}
		}
	}
}
