package auth

import (
	"net/http"

	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/glue/routing"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers"
	"github.com/gin-gonic/gin"

	"go.uber.org/zap"
)

func Init(
	group *gin.RouterGroup,
	log zap.Logger,
	authHandler handlers.Auth,
) {

	weatherRoutes := []routing.Route{
		{
			Method:     http.MethodPost,
			Path:       "/user",
			Handler:    authHandler.RegisterUser,
			Middleware: []gin.HandlerFunc{},
		}, {
			Method:     http.MethodPost,
			Path:       "/user/login",
			Handler:    authHandler.Login,
			Middleware: []gin.HandlerFunc{},
		},
	}
	routing.RegisterRoute(group, weatherRoutes, log)
}
