package impl

import (
	taskDb "base/internal/outbound/db/task"
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"

	"gorm.io/gorm"
)

type repository struct {
	Db     *gorm.DB
	Config pkgConfig.Config
}

func New(dbList model.Database, cfg pkgConfig.Config) taskDb.Repository {
	return &repository{
		Db:     dbList.Template,
		Config: cfg,
	}
}
