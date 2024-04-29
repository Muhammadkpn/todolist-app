package impl

import (
	healthCheckUsecase "base/internal/usecase/healthCheck"
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"
)

type usecase struct {
	Db  model.Database
	cfg pkgConfig.Config
}

func New(dbList model.Database, cfg pkgConfig.Config) healthCheckUsecase.Usecase {
	return &usecase{
		Db:  dbList,
		cfg: cfg,
	}
}
