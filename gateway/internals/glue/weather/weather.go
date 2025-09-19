package weather

import (
	"net/http"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"

	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/glue/routing"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init(
	group *gin.RouterGroup,
	log zap.Logger,
	weatherHandler handlers.Weather,
	grpcClient pb.AuthServiceClient,

) {

	weatherRoutes := []routing.Route{
		{
			Method:  http.MethodPost,
			Path:    "/weather",
			Handler: weatherHandler.HandleGetWeather,
			Middleware: []gin.HandlerFunc{
				middleware.Auth(grpcClient, log),
			},
		},
	}
	routing.RegisterRoute(group, weatherRoutes, log)
}
