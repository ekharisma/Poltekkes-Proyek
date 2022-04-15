package constant

const (
	// DB Constant
	DBUsername = "ekky"
	DBPassword = "password"
	DBHost     = "localhost"
	DBName     = "poltekkes-db"
	DBPort     = 5432
	//MQTT Constant
	Broker   = "broker.hivemq.com"
	MqttPort = 1883
	//SMTP
	SMTPHost       = "smtp.gmail.com"
	SMTPPort       = 587
	SMTPSenderName = "Poltekkes Temperature Observatory"
	SMTPEmail      = "backendprogrammer43@gmail.com"
	SMTPPassword   = "CahayaMei123"
	//Telegram
	TelegramChannel = "vaccinebox"
	TelegramURI     = "https://api.telegram.org/bot%v/sendMessage?chat_id=@%v&text=%v"
	TelegramToken   = "5137342139:AAE1OAGP0VxNyQTQV9tdT8hH8zkBTpx4WfE"
)
