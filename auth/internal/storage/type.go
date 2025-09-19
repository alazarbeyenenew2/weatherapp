package storage

import (
	"context"

	"github.com/alazarbeyeneazu/weatherapp/common/models"
)

type User interface {
	SaveUser(ctx context.Context, usr models.User) (models.User, error)
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}
