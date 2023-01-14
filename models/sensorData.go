package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SensorDataIn struct {
	Temperature float32 `json:"temp" xml:"temp" form:"temp" query:"temp"`
	Humidity    float32 `json:"hum" xml:"hum" form:"hum" query:"hum"`
}

type SensorDataOut struct {
	Temperature float32   `json:"temp" xml:"temp" form:"temp" query:"temp"`
	Humidity    float32   `json:"hum" xml:"hum" form:"hum" query:"hum"`
	Time        time.Time `json:"time"`
}

type SensorDataMongo struct {
	Temperature float32
	Humidity    float32
	Time        primitive.DateTime
}
