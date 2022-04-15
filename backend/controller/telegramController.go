package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ekharisma/poltekkes-webservice/constant"
	"github.com/ekharisma/poltekkes-webservice/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramController struct {
	telegramClient  *tgbotapi.BotAPI
	lastMessageSent time.Time
	interval        time.Duration
}

type ITelegramController interface {
	ProcessMessageRequest(temperature entity.Temperature)
	SendMessage(temperature entity.Temperature)
}

func CreateTelegramController(telegramClient *tgbotapi.BotAPI) ITelegramController {
	return &TelegramController{
		telegramClient: telegramClient,
		interval:       1800 * time.Second,
	}
}

func constructMessage(temperature entity.Temperature) string {
	return fmt.Sprintf(`Pada %v, suhu diluar batas normal dengan suhu : %v celcius`, temperature.Timestamp, temperature.Temperature)
}

func (c *TelegramController) ProcessMessageRequest(temperature entity.Temperature) {
	if !c.isRateLimited() {
		go c.SendMessage(temperature)
	}
}

func (c *TelegramController) SendMessage(temperature entity.Temperature) {
	message := constructMessage(temperature)
	url := fmt.Sprintf(constant.TelegramURI, constant.TelegramToken, constant.TelegramChannel, message)
	fmt.Println(url)
	_, err := http.Get(url)
	if err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
}

func (c *TelegramController) isRateLimited() bool {
	if (c.lastMessageSent == time.Time{}) {
		c.lastMessageSent = time.Now()
		return false
	} else if time.Since(c.lastMessageSent) > c.interval {
		fmt.Println("Melebihi interval")
		c.lastMessageSent = time.Now()
		return false
	}
	return true
}
