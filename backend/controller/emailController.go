package controller

import (
	"fmt"
	"github.com/ekharisma/poltekkes-webservice/constant"
	"github.com/ekharisma/poltekkes-webservice/entity"
	"github.com/ekharisma/poltekkes-webservice/model"
	"log"
	"net/smtp"
)

type EmailController struct {
	userModel model.IUserModel
}

type IEmailController interface {
	SendEmail(temperature entity.Temperature)
	getEmailToSend() []string
}

func CreateEmailController(userModel model.IUserModel) IEmailController {
	return &EmailController{
		userModel: userModel,
	}
}

func (e EmailController) getEmailToSend() (emails []string) {
	users, err := e.userModel.GetAll()
	if err != nil {
		return nil
	}
	for _, user := range users {
		emails = append(emails, user.Email)
	}
	return emails
}

func (e EmailController) SendEmail(temperature entity.Temperature) {
	emails := e.getEmailToSend()
	subject := "Peringatan : Suhu Diluar batas normal"
	message := fmt.Sprintf(`
		Peringatan suhu yang diamati diluar batas. Berikut data suhu dan waktu pengambilan data.
		Suhu : %v Celcius,
		Waktu : %v
		Harap segera melakukan pengecekan pada alat. Terima kasih
	`, temperature.Temperature, temperature.Timestamp)
	body := fmt.Sprintf(`
		From: %v
		To : %v
		Subject : %v
		%v
	`, constant.SMTPSenderName, emails, subject, message)
	auth := smtp.PlainAuth("", constant.SMTPEmail, constant.SMTPPassword, constant.SMTPHost)
	smtpAddress := fmt.Sprintf("%s:%d", constant.SMTPHost, constant.SMTPPort)
	err := smtp.SendMail(smtpAddress, auth, constant.SMTPEmail, emails, []byte(body))
	if err != nil {
		log.Panicln("Error. Reason : ", err.Error())
	}
}
