package impl

import (
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"

	taskCache "base/internal/outbound/cache/task"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	Cache  *redis.Client
	Config pkgConfig.Config
}

func New(cacheList model.Redis, cfg pkgConfig.Config) taskCache.Repository {
	return &repository{
		Cache:  cacheList.Main,
		Config: cfg,
	}
}
