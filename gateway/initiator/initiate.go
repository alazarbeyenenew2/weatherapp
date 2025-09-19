package initiator

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alazarbeyeneazu/weatherapp/gateway/internals/handlers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Initiate() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("unable to start logger")
	}
	configName := "config"
	if os.Getenv("CONFIG_NAME") != "" {
		configName = os.Getenv("CONFIG_NAME")
	}
	initConfig(configName, "config", logger)

	//initializing grpc client
	grpc := initGRPC()

	//initializing handler
	logger.Info("initializing handler")
	handler := initHandler(grpc, *logger)
	logger.Info("handler initialized")

	//initializing server
	server := gin.New()
	server.Use(middleware.GinLogger(*logger))
	server.Use(middleware.CORS())
	grp := server.Group("/api")
	initRoute(grp, handler, *logger, grpc.auth)

	//initializing http server
	srv := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", viper.GetString("app.host"), viper.GetInt("app.port")),
		Handler:           server,
		ReadHeaderTimeout: viper.GetDuration("app.timeout"),
	}

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT)
		<-sigint
		log.Fatal("HTTP server Shutdown")

	}()
	logger.Info(fmt.Sprintf("http server listening on port : %d", viper.GetInt("app.port")))
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal(fmt.Sprintf("Could not start HTTP server: %s", err))
	}

}
