package data

import (
	"context"
	c "github.com/PLDao/gin-frame/config"
	log "github.com/PLDao/gin-frame/internal/pkg/logger"
	"github.com/go-redis/redis/v8"
)

func initRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     c.Config.Redis.Host + ":" + c.Config.Redis.Port,
		Password: c.Config.Redis.Password,
		DB:       c.Config.Redis.Database,
	})
	var ctx = context.Background()
	_, err := Rdb.Ping(ctx).Result()

	if err != nil {
		panic("Redis connection failed：" + err.Error())
	}
	log.Logger.Sugar().Info(" Redis connection successful")
}
