package activedirectory

import (
	interfaces "base/internal/outbound/interface"
	pkgConfig "base/pkg/config"
)

type repository struct {
	cfg pkgConfig.Config
}

func New(cfg pkgConfig.Config) interfaces.Authentication {
	return &repository{
		cfg: cfg,
	}
}
