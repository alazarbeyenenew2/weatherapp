package weather

import (
	"net/http"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"

	"github.com/alazarbeyeneazu/weatherapp/common"
	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type weatherHandler struct {
	client pb.WeatherServiceClient
	log    *zap.Logger
}

func Init(client pb.WeatherServiceClient, log *zap.Logger) handlers.Weather {
	return &weatherHandler{
		client: client,
		log:    log,
	}
}
func (w *weatherHandler) HandleGetWeather(c *gin.Context) {
	var weatherRequest models.WeatherRequest
	if err := c.ShouldBind(&weatherRequest); err != nil {
		w.log.Warn("unable to bind request to models.WeatherRequest", zap.Error(err))
		common.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	weather, err := w.client.GetWeather(c, &pb.WeatherRequest{
		Location: weatherRequest.Location,
		Datetime: weatherRequest.DateTime,
	})
	if err != nil {
		rStatus := status.Convert(err)
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(c, http.StatusInternalServerError, rStatus.Message())
			return
		}
		common.WriteError(c, http.StatusBadRequest, rStatus.Message())
		return
	}
	common.WriteJSON(c, http.StatusOK, weather)
}
