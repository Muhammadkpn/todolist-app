package injection

import (
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg pkgConfig.Config) (cache model.Redis, err error) {
	opt, err := redis.ParseURL(cfg.Redis.Main.GenerateConnectionString())
	if err != nil {
		return
	}

	main := redis.NewClient(opt)

	_, err = main.Ping(context.Background()).Result()
	if err != nil {
		return
	}

	cache.Main = main

	return
}
