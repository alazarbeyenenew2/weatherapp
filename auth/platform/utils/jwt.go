package utils

import (
	"fmt"
	"time"

	"github.com/alazarbeyeneazu/weatherapp/common/models"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func GenerateJWT(user models.User, secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"id":         user.ID.Hex(),
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"country":    user.Country,
		"city":       user.City,
		"email":      user.Email,
		"exp":        time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
func ValidateJWT(tokenString, secretKey string, log *zap.Logger) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			log.Error("unexpected signing method", zap.Error(err), zap.Any("request", tokenString))
			return nil, err
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		log.Error("unable to parse token ", zap.Error(err), zap.Any("request", tokenString))
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, _ := primitive.ObjectIDFromHex(claims["id"].(string))
		user := &models.User{
			ID:        id,
			FirstName: claims["first_name"].(string),
			LastName:  claims["last_name"].(string),
			Country:   claims["country"].(string),
			City:      claims["city"].(string),
			Email:     claims["email"].(string),
		}
		return user, nil
	}
	err = fmt.Errorf("invalid token")
	log.Error("invalid token", zap.Any("request", tokenString))
	return nil, err
}
