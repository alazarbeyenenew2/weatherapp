package initiator

import (
	"github.com/alazarbeyeneazu/weatherapp/auth/internal/storage"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type persistence struct {
	userDB storage.User
}

func initPersistence(db *mongo.Client, log *zap.Logger) persistence {
	return persistence{
		userDB: storage.Init(getCollection(db, "users"), log),
	}
}
func getCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(viper.GetString("database.db_name")).Collection(collectionName)
	return collection
}
