package initiator

import (
	"log"

	pb "github.com/alazarbeyeneazu/weatherapp/common/api"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPC struct {
	weather pb.WeatherServiceClient
	auth    pb.AuthServiceClient
}
type gRPCConnections struct {
	WeatherConnection *grpc.ClientConn
	AuthConnection    *grpc.ClientConn
}

func getConnections() *gRPCConnections {
	weatherGRPCClient, err := grpc.NewClient(viper.GetString("weather.grpc"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("unable to get grpc weather client connection")
	}
	authGRPCClient, err := grpc.NewClient(viper.GetString("auth.grpc"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("unable to get grpc auth client connection")
	}
	return &gRPCConnections{
		WeatherConnection: weatherGRPCClient,
		AuthConnection:    authGRPCClient,
	}
}
func initGRPC() gRPC {
	connections := getConnections()
	return gRPC{
		weather: pb.NewWeatherServiceClient(connections.WeatherConnection),
		auth:    pb.NewAuthServiceClient(connections.AuthConnection),
	}
}
