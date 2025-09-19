package initiator

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func initDB(url string, log *zap.Logger) *mongo.Client {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal("unable to connect to mongodb server")
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("unable to ping mongodb server")
	}
	return client
}
