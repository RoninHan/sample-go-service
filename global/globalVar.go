package global

import (
	"context"
	"sample-go-service/config"
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"go.uber.org/zap"
)

var (
	Settings config.ServerConfig
	//globalVar.go
	DB    *gorm.DB
	Redis *redis.Client
)

var (
	Lg *zap.Logger
)
var (
	Trans ut.Translator
)

func GetCtx() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	return ctx
}
