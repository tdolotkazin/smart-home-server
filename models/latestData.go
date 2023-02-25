package models

import "time"

type AllData struct {
	Temperature           float32   `json:"temp"`
	Humidity              float32   `json:"hum"`
	BaseLastUpdatedTime   time.Time `json:"baseLastUpdatedTime"`
	WaterTemperature      float32   `json:"waterTemp"`
	IsRunning             bool      `json:"isRunning"`
	BoilerLastUpdatedTime time.Time `json:"boilerLastUpdatedTime"`
}
