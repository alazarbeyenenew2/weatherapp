package middleware

import (
	"net/http"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Auth(client pb.AuthServiceClient, log zap.Logger) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		tokenString := authHeader[7:]
		usr, err := client.ValidateToken(ctx, &pb.LoginResponse{Token: tokenString})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Set("email", usr.Email)
		ctx.Next()
	}
}
