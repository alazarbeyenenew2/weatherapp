package handler

import (
	"context"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"

	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"github.com/alazarbeyeneazu/weatherapp/weather/internal/module"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedWeatherServiceServer
	service module.WeatherService
}

func NewGRPCHandler(grpcServer *grpc.Server, service module.WeatherService) {
	handler := &grpcHandler{
		service: service,
	}
	pb.RegisterWeatherServiceServer(grpcServer, handler)
}
func (g *grpcHandler) GetWeather(ctx context.Context, rq *pb.WeatherRequest) (*pb.Weather, error) {
	rqj := models.WeatherRequest{
		Location: rq.Location,
		DateTime: rq.Datetime,
	}
	respons, err := g.service.GetWeather(ctx, rqj)
	if err != nil {
		return &pb.Weather{}, err
	}
	//
	resp := respons.Days[0]
	daily := &pb.WeatherData{
		Datetime:  resp.Datetime,
		Tempmin:   resp.Tempmin,
		Tempmax:   resp.Tempmax,
		Humidity:  resp.Humidity,
		Precip:    resp.Precip,
		Snow:      resp.Snow,
		Snowdepth: resp.Snowdepth,
		Windspeed: resp.Windspeed,
	}
	hourly := []*pb.WeatherData{}
	for _, weatherHourly := range resp.Hours {
		hourly = append(hourly, &pb.WeatherData{
			Datetime:  weatherHourly.Datetime,
			Tempmin:   weatherHourly.Tempmin,
			Tempmax:   weatherHourly.Tempmax,
			Humidity:  weatherHourly.Humidity,
			Precip:    weatherHourly.Precip,
			Snow:      weatherHourly.Snow,
			Temp:      weatherHourly.Temp,
			Snowdepth: weatherHourly.Snowdepth,
			Windspeed: weatherHourly.Windspeed,
		})
	}

	return &pb.Weather{
		Day:    daily,
		Hourly: hourly,
	}, nil
}
