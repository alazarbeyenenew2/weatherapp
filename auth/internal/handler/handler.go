package handler

import (
	"context"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"

	"github.com/alazarbeyeneazu/weatherapp/auth/internal/module"
	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"google.golang.org/grpc"
)

type gRPCHandler struct {
	pb.UnimplementedAuthServiceServer
	service module.Auth
}

func NewGRPCHandler(gRPCServer *grpc.Server, service module.Auth) {
	handler := &gRPCHandler{
		service: service,
	}
	pb.RegisterAuthServiceServer(gRPCServer, handler)
}
func (a *gRPCHandler) RegisterUser(ctx context.Context, rq *pb.RegisterRequest) (*pb.LoginResponse, error) {
	userRequest := models.User{
		FirstName: rq.FirstName,
		LastName:  rq.LastName,
		Country:   rq.Country,
		City:      rq.City,
		Email:     rq.Email,
		Password:  rq.Password,
	}

	resp, err := a.service.RegisterUser(ctx, userRequest)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: resp.Token,
	}, nil
}
func (a *gRPCHandler) LoginUser(ctx context.Context, rq *pb.LoginRequest) (*pb.LoginResponse, error) {
	loginRequest := models.UserLoginRequest{
		Email:    rq.Email,
		Password: rq.Password,
	}
	resp, err := a.service.Login(ctx, loginRequest)
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Token: resp.Token,
	}, nil
}
func (a *gRPCHandler) ValidateToken(ctx context.Context, rq *pb.LoginResponse) (*pb.RegisterRequest, error) {
	req := models.UserLoginResponse{
		Token: rq.Token,
	}
	resp, err := a.service.VerifyUser(ctx, req)

	if err != nil {
		return nil, err
	}
	return &pb.RegisterRequest{
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		Country:   resp.Country,
		City:      resp.City,
		Email:     resp.Email,
	}, nil
}
