package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func IsDataExist(ctx context.Context, filter bson.M, log *zap.Logger, collection *mongo.Collection) bool {
	err := collection.FindOne(ctx, filter).Err()

	if err != nil && err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		log.Error("unable to check data existance", zap.Error(err))
		return false
	}

	return true

}
