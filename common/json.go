package common

import (
	"github.com/gin-gonic/gin"
)

func WriteJSON(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func WriteError(c *gin.Context, status int, message string) {
	WriteJSON(c, status, map[string]string{"error": message})
}
