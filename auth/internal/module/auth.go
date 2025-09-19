package module

import (
	"context"
	"fmt"

	"github.com/alazarbeyeneazu/weatherapp/auth/internal/storage"
	"github.com/alazarbeyeneazu/weatherapp/auth/platform/utils"
	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type auth struct {
	userStorage storage.User
	log         *zap.Logger
	authSec     string
}

func Init(userStorage storage.User, authSec string, log *zap.Logger) Auth {
	return &auth{
		userStorage: userStorage,
		log:         log,
		authSec:     authSec,
	}
}
func (a *auth) RegisterUser(ctx context.Context, user models.User) (models.UserLoginResponse, error) {
	if err := user.Validate(); err != nil {
		a.log.Warn(err.Error(), zap.Error(err), zap.Any("request", user))
		return models.UserLoginResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	//hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		a.log.Warn(err.Error(), zap.Error(err), zap.Any("request", user))
		return models.UserLoginResponse{}, status.Errorf(codes.Internal, err.Error())
	}
	user.Password = hashedPassword

	// saving user to database
	userResp, err := a.userStorage.SaveUser(ctx, user)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	token, err := utils.GenerateJWT(userResp, a.authSec)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	return models.UserLoginResponse{Token: token}, nil

}
func (a *auth) Login(ctx context.Context, user models.UserLoginRequest) (models.UserLoginResponse, error) {
	if err := user.Validate(); err != nil {
		a.log.Warn(err.Error(), zap.Error(err), zap.Any("request", user))
		return models.UserLoginResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	//getUserByEmail
	userResp, err := a.userStorage.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	// validate password
	if !utils.ComparePassword(userResp.Password, user.Password) {
		err = fmt.Errorf("invalid username or password")
		a.log.Warn(err.Error(), zap.Error(err), zap.Any("request", user))
		return models.UserLoginResponse{}, status.Errorf(codes.InvalidArgument, err.Error())
	}
	//remove password from response
	userResp.Password = ""
	//generate jwtToken

	token, err := utils.GenerateJWT(userResp, a.authSec)
	if err != nil {
		a.log.Warn(err.Error(), zap.Error(err), zap.Any("request", user))
		return models.UserLoginResponse{}, err
	}

	return models.UserLoginResponse{
		Token: token,
	}, nil

}

func (a *auth) VerifyUser(ctx context.Context, rq models.UserLoginResponse) (models.User, error) {
	user, err := utils.ValidateJWT(rq.Token, a.authSec, a.log)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}
