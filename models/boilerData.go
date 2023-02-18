package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BoilerDataIn struct {
	WaterTemperature float32 `json:"waterTemp" xml:"waterTemp" form:"waterTemp" query:"waterTemp"`
	IsRunning        bool    `json:"isRunning" xml:"isRunning" form:"isRunning" query:"isRunning"`
}

type BoilerDataOut struct {
	WaterTemperature float32   `json:"waterTemp" xml:"waterTemp" form:"waterTemp" query:"waterTemp"`
	IsRunning        bool      `json:"isRunning" xml:"isRunning" form:"isRunning" query:"isRunning"`
	Time             time.Time `json:"time"`
}

type BoilerDataMongo struct {
	WaterTemperature float32
	IsRunning        bool
	Time             primitive.DateTime
}
