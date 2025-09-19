package storage

import (
	"context"

	"github.com/alazarbeyeneazu/weatherapp/auth/platform/utils"
	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type user struct {
	userCollection *mongo.Collection
	log            *zap.Logger
}

func Init(userCollection *mongo.Collection, log *zap.Logger) User {
	return &user{
		userCollection: userCollection,
		log:            log,
	}
}
func (u *user) SaveUser(ctx context.Context, usr models.User) (models.User, error) {
	if utils.IsDataExist(ctx, bson.M{"email": usr.Email}, u.log, u.userCollection) {
		u.log.Warn("user already saved ", zap.Any("request", usr))
		return models.User{}, status.Errorf(codes.InvalidArgument, "user already saved")
	}
	id := primitive.NewObjectID()
	usr.ID = id
	_, err := u.userCollection.InsertOne(ctx, usr)
	if err != nil {
		u.log.Error("unable to save user", zap.Error(err), zap.Any("user", usr))
		return models.User{}, status.Errorf(codes.Internal, err.Error())
	}
	return usr, nil
}
func (u *user) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var userMD models.User
	if !utils.IsDataExist(ctx, bson.M{"email": email}, u.log, u.userCollection) {
		u.log.Warn("user not exist ", zap.Any("request", email))
		return models.User{}, status.Errorf(codes.InvalidArgument, "user not exist")
	}
	if err := u.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&userMD); err != nil {
		u.log.Error("error while getting user", zap.Any("request", email))
		return models.User{}, status.Errorf(codes.Internal, "error while getting user")
	}
	return userMD, nil

}
