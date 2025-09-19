package initiator

import (
	md "github.com/alazarbeyeneazu/weatherapp/auth/internal/module"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type module struct {
	authModule md.Auth
}

func initModule(persistenceDB persistence, log *zap.Logger) module {
	return module{
		authModule: md.Init(persistenceDB.userDB, viper.GetString("auth.jwtsecret"), log),
	}
}
