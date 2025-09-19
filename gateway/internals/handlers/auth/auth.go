package auth

import (
	"net/http"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"
	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers"

	"github.com/alazarbeyeneazu/weatherapp/common"
	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type auth struct {
	client pb.AuthServiceClient
	log    *zap.Logger
}

func Init(client pb.AuthServiceClient, log *zap.Logger) handlers.Auth {
	return &auth{
		client: client,
		log:    log,
	}
}
func (a *auth) RegisterUser(c *gin.Context) {
	var userRr models.User
	if err := c.ShouldBind(&userRr); err != nil {
		a.log.Warn("unable to bind request to models.UserRegistrationRequest", zap.Error(err))
		common.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := a.client.RegisterUser(c, &pb.RegisterRequest{
		FirstName: userRr.FirstName,
		LastName:  userRr.LastName,
		Country:   userRr.Country,
		City:      userRr.City,
		Email:     userRr.Email,
		Password:  userRr.Password,
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
	common.WriteJSON(c, http.StatusOK, resp)

}

func (a *auth) Login(c *gin.Context) {
	var userRr models.UserLoginRequest
	if err := c.ShouldBind(&userRr); err != nil {
		a.log.Warn("unable to bind request to models.UserLoginRequest", zap.Error(err))
		common.WriteError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := a.client.LoginUser(c, &pb.LoginRequest{
		Email:    userRr.Email,
		Password: userRr.Password,
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
	common.WriteJSON(c, http.StatusOK, resp)

}
