package platform

import (
	"context"

	"github.com/alazarbeyeneazu/weatherapp/common/models"
)

type WeatherAPI interface {
	GetWeather(ctx context.Context, rq models.WeatherRequest, response *models.WeatherResponse) error
}
