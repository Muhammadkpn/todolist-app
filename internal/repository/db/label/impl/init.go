package impl

import (
	labelRepository "base/internal/repository/db/label"
	pkgConfig "base/pkg/config"

	"gorm.io/gorm"
)

type repository struct {
	DbGorm *gorm.DB
	Config pkgConfig.Config
}

func New(dbGorm *gorm.DB, cfg pkgConfig.Config) labelRepository.Repository {
	return &repository{
		DbGorm: dbGorm,
		Config: cfg,
	}
}
