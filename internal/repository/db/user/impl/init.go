package impl

import (
	userRepository "base/internal/repository/db/user"
	pkgConfig "base/pkg/config"

	"gorm.io/gorm"
)

type repository struct {
	DbGorm *gorm.DB
	Config pkgConfig.Config
}

func New(dbGorm *gorm.DB, cfg pkgConfig.Config) userRepository.Repository {
	return &repository{
		DbGorm: dbGorm,
		Config: cfg,
	}
}
