package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/ekharisma/sensor-emu/sensoremu"
)

type Payload struct {
	Timestamp   int64
	Temperature float32
}

func main() {
	rand.Seed(time.Now().Unix())
	client := sensoremu.InitBroker()
	for {
		data := sensoremu.ProduceData()
		time.Sleep(2 * time.Second)
		fmt.Println("Data produced : ", data)
		payload := &Payload{
			Timestamp:   time.Now().UnixMilli(),
			Temperature: data,
		}
		message, _ := json.Marshal(payload)
		client.Publish("/emulator/data", 2, false, message)
	}
}
