package impl

import (
	healthCheckUsecase "base/internal/usecase/healthCheck"
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"
)

type usecase struct {
	Db    model.Database
	Redis model.Redis
	cfg   pkgConfig.Config
}

func New(dbList model.Database, redisList model.Redis, cfg pkgConfig.Config) healthCheckUsecase.Usecase {
	return &usecase{
		Db:    dbList,
		Redis: redisList,
		cfg:   cfg,
	}
}
