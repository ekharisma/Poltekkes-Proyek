package sensoremu

import "math/rand"

func ProduceData() float32 {
	return rand.Float32() * 10
}

func IsDataSafe(data float32) bool {
	if data < 2.0 || data > 8 {
		return false
	}
	return true
}
