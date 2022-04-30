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
	evacuationCode  int
}

type ITelegramController interface {
	DoorOpened()
	OutOfRangeTemperature(temperature entity.Temperature)
	Evacuation() int
	EarlyWarningProcessor(temperature entity.Temperature)
	SendMessage(message string)
}

func CreateTelegramController(telegramClient *tgbotapi.BotAPI) ITelegramController {
	return &TelegramController{
		telegramClient: telegramClient,
		interval:       1800 * time.Second,
		evacuationCode: 1,
	}
}

func constructEarlyWarningMessage(temperature entity.Temperature) string {
	return fmt.Sprintf(`Early Warning : Pada %v, suhu diluar batas normal dengan suhu : %v celcius`, temperature.Timestamp, temperature.Temperature)
}

func constructOutOfRangeMessage(temperture entity.Temperature) string {
	return fmt.Sprintf(`Pada %v, suhu telah out of range. Suhu terpantau %v celcius`, temperture.Timestamp, temperture.Temperature)
}

func constructDoorOpenedMessage() string {
	return "Pintu terbuka"
}

func constructEvacutionMessage(code int) string {
	return fmt.Sprintf(`Peringatan evakuasi ke-%v. Segera lakukan prosedur evakuasi`, code)
}

func (c *TelegramController) OutOfRangeTemperature(temperature entity.Temperature) {
	message := constructOutOfRangeMessage(temperature)
	go c.SendMessage(message)
}

func (c *TelegramController) DoorOpened() {
	message := constructDoorOpenedMessage()
	go c.SendMessage(message)
}

func (c *TelegramController) Evacuation() int {
	message := constructEvacutionMessage(c.evacuationCode)
	fmt.Println("Code : ", c.evacuationCode)
	if c.evacuationCode >= 4 {
		message = `Peringatan evakuasi telah habis`
		c.evacuationCode = 1
	}
	c.evacuationCode += 1
	go c.SendMessage(message)
	return c.evacuationCode
}

func (c *TelegramController) EarlyWarningProcessor(temperature entity.Temperature) {
	message := constructEarlyWarningMessage(temperature)
	go c.SendMessage(message)
}

func (c *TelegramController) SendMessage(message string) {
	url := fmt.Sprintf(constant.TelegramURI, constant.TelegramToken, constant.TelegramChannel, message)
	_, err := http.Get(url)
	if err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
}

//func (c *TelegramController) isRateLimited() bool {
//	if (c.lastMessageSent == time.Time{}) {
//		c.lastMessageSent = time.Now()
//		return false
//	} else if time.Since(c.lastMessageSent) > c.interval {
//		c.lastMessageSent = time.Now()
//		return false
//	}
//	return true
//}
