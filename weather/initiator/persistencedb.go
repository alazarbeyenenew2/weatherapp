package initiator

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type persistence struct{}

func initPersistence(db *mongo.Client, log *zap.Logger) persistence {
	return persistence{}
}
