package service

import (
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttClient mqtt.Client
var mqttClientSingleton sync.Once

func CreateMqttClient() mqtt.Client {
	mqttClientSingleton.Do(func() {
		opts := mqtt.NewClientOptions()
		opts.AddBroker()
	})
}
