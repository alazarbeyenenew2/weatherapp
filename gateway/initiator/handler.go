package initiator

import (
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers/auth"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers/weather"
	"go.uber.org/zap"
)

type handler struct {
	weather handlers.Weather
	auth    handlers.Auth
}

func initHandler(gRPC gRPC, log zap.Logger) handler {
	return handler{
		weather: weather.Init(gRPC.weather, &log),
		auth:    auth.Init(gRPC.auth, &log),
	}
}
