package impl

import (
	taskRepository "base/internal/repository/db/task"
	pkgConfig "base/pkg/config"

	"gorm.io/gorm"
)

type repository struct {
	DbGorm *gorm.DB
	Config pkgConfig.Config
}

func New(dbGorm *gorm.DB, cfg pkgConfig.Config) taskRepository.Repository {
	return &repository{
		DbGorm: dbGorm,
		Config: cfg,
	}
}
