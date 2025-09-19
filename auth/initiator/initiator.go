package initiator

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/alazarbeyeneazu/weatherapp/auth/internal/handler"
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

	//initializing module
	logger.Info("initializing module")
	module := initModule(persistancedb, logger)
	logger.Info("modules initialized")

	// initializing grpc server
	grpcServer := grpc.NewServer()
	logger.Info(fmt.Sprintf("auth service start at  %s", viper.GetString("grpcAddr")))
	listen, err := net.Listen("tcp", viper.GetString("grpcAddr"))
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()
	handler.NewGRPCHandler(grpcServer, module.authModule)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
