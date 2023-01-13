package models

type SensorData struct {
	Temperature float32 `json:"temp" xml:"temp" form:"temp" query:"temp"`
	Humidity    float32 `json:"hum" xml:"hum" form:"hum" query:"hum"`
}
