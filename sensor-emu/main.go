package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/ekharisma/sensor-emu/sensoremu"
)

type Payload struct {
	Timestamp   time.Time
	Temperature float32
}

func main() {
	rand.Seed(time.Now().Unix())
	client := sensoremu.InitBroker()
	for {
		data := sensoremu.ProduceData()
		time.Sleep(2 * time.Second)
		payload := &Payload{
			Timestamp:   time.Now(),
			Temperature: data,
		}
		message, _ := json.Marshal(payload)
		fmt.Printf("Data produced : %v \n", payload)
		client.Publish("/emulator/data", 2, false, message)
	}
}
