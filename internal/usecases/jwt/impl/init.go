package impl

import (
	uJwt "base/internal/usecases/jwt"
	pkgConfig "base/pkg/config"
)

type usecase struct {
	cfg pkgConfig.Config
}

func New(cfg pkgConfig.Config) uJwt.Usecase {
	return &usecase{
		cfg: cfg,
	}
}
