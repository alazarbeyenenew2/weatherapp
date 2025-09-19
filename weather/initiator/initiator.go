package initiator

import (
	"log"
	"net"
	"os"

	"github.com/alazarbeyeneazu/weatherapp/weather/internal/handler"
	weatherapi "github.com/alazarbeyeneazu/weatherapp/weather/platform/weatherAPI"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Init() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to start logger")
	}
	configName := "config"
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}
	// Initializing config
	logger.Info("initializing config")
	initConfig(configName, "config", logger)
	logger.Info("config initialization complited")

	//initializing database
	logger.Info("initializing database")
	db := initDB(viper.GetString("database.url"), logger)
	logger.Info("database initialized")

	//initializing persistence
	logger.Info("initializing persistence ")
	persistancedb := initPersistence(db, logger)
	logger.Info("peristence layer initialized")

	//initializing platform layer
	logger.Info("initializing platform ")
	baseurl := viper.GetString("visualcrossing.base_url")
	weatherapi := weatherapi.InitVisualcrossing(baseurl, logger)
	logger.Info("platform initialized")

	//initializing module
	logger.Info("initializing module")
	module := initModule(persistancedb, weatherapi, logger)
	logger.Info("modules initialized")

	// initializing grpc server
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", viper.GetString("grpcAddr"))
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	handler.NewGRPCHandler(grpcServer, module.weatherModule)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
