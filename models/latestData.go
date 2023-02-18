package models

import "time"

type LatestData struct {
	Temperature      float32   `json:"temp"`
	Humidity         float32   `json:"hum"`
	WaterTemperature float32   `json:"waterTemp"`
	IsRunning        bool      `json:"isRunning"`
	Time             time.Time `json:"time"`
}
