package module

import (
	"context"

	"github.com/alazarbeyeneazu/weatherapp/common/models"
)

type Auth interface {
	RegisterUser(ctx context.Context, user models.User) (models.UserLoginResponse, error)
	Login(ctx context.Context, user models.UserLoginRequest) (models.UserLoginResponse, error)
	VerifyUser(ctx context.Context, rq models.UserLoginResponse) (models.User, error)
}
