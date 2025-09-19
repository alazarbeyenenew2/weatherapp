package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type WeatherRequest struct {
	Location string `json:"location" bson:"location"`
	DateTime string `json:"datetime" bson:"datetime"`
}
type Weather struct {
	Datetime  string    `json:"datetime" bson:"datetime"`
	Tempmin   float32   `json:"tempmin" bson:"tempmin"`
	Tempmax   float32   `json:"tempmax" bson:"tempmax"`
	Humidity  float32   `json:"humidity" bson:"humidity"`
	Precip    float32   `json:"precip" bson:"precip"`
	Snow      float32   `json:"snow" bson:"snow"`
	Snowdepth float32   `json:"snowdepth" bson:"snowdepth"`
	Windspeed float32   `json:"windspeed" bson:"windspeed"`
	Temp      float32   `json:"temp" bson:"temp"`
	Hours     []Weather `json:"hours" bson:"hour"`
}
type WeatherResponse struct {
	Days []Weather `json:"days" bson:"days"`
}

func (w WeatherRequest) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.DateTime, validation.Required.Error("datetime field required")),
		validation.Field(&w.Location, validation.Required.Error("location field required")),
	)
}
