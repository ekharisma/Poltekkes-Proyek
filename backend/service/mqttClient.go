package service

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client
var mqttClientSingleton sync.Once

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func CreateMqttClient(broker string, port uint) {
	mqttClientSingleton.Do(func() {
		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
		opts.SetClientID("Poltekkes-TA")
		opts.OnConnect = connectHandler
		opts.OnConnectionLost = connectLostHandler
		opts.SetDefaultPublishHandler(messagePubHandler)
		mqttClient = mqtt.NewClient(opts)
		if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	})
}

func GetMqttClient() mqtt.Client {
	return mqttClient
}
