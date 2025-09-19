package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Country   string             `json:"country" bson:"country"`
	City      string             `json:"city" bson:"city"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
}
type UserLoginResponse struct {
	Token string `json:"token"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first_Name required")),
		validation.Field(&u.LastName, validation.Required.Error("last_name required")),
		validation.Field(&u.Country, validation.Required.Error("country required")),
		validation.Field(&u.City, validation.Required.Error("city required")),
		validation.Field(&u.Email, validation.Required.Error("email required")),
		validation.Field(&u.Password, validation.Required.Error("password required")),
	)
}

type UserLoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (u UserLoginRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required.Error("email required")),
		validation.Field(&u.Password, validation.Required.Error("password required")),
	)
}
