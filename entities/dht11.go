package entities

import "time"

type DHT11 struct {
	ID          int       `json:"id"`
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Is_UDP      bool      `json:"is_UDP"`
	RegisterAt  time.Time `json:"registerAt"`
}

func NewDHT11(temperature, humidity float32, isUDP bool) *DHT11 {
	date := time.Now()

	return &DHT11{
		Temperature: temperature,
		Humidity:    humidity,
		Is_UDP:      isUDP,
		RegisterAt:  date,
	}
}
