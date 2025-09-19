package weatherapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"github.com/alazarbeyeneazu/weatherapp/weather/platform"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type visualcrossing struct {
	baseUrl string
	log     *zap.Logger
}

func InitVisualcrossing(baseUrl string, log *zap.Logger) platform.WeatherAPI {
	return &visualcrossing{
		baseUrl: baseUrl,
		log:     log,
	}
}

func (v *visualcrossing) GetWeather(ctx context.Context, rq models.WeatherRequest, response *models.WeatherResponse) error {
	resp, err := http.Get(fmt.Sprintf(v.baseUrl, rq.Location, rq.DateTime))
	if err != nil {
		v.log.Error("unable to get weather information ", zap.Error(err), zap.Any("request ", rq))
		return status.Errorf(codes.Internal, err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		v.log.Error("got unexpected status return", zap.Error(err), zap.Any("request ", rq))
		return status.Errorf(codes.Internal, "unexpected status code")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		v.log.Error("error while reading response body", zap.Error(err), zap.Any("request", rq))
		return status.Error(codes.Internal, fmt.Sprintf("error reading response body: %v", err))
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		v.log.Error("error unmarshalling JSON", zap.Error(err), zap.Any("request", rq))
		return status.Errorf(codes.Internal, err.Error())
	}

	return nil
}
