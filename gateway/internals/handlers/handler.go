package handlers

import (
	"github.com/gin-gonic/gin"
)

type Weather interface {
	HandleGetWeather(c *gin.Context)
}
type Auth interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
}
