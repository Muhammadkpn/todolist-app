package impl

import (
	exampleNonDiRepository "base/internal/outbound/db/exampleNonDi"
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"

	"gorm.io/gorm"
)

type repository struct {
	Db     *gorm.DB
	Config pkgConfig.Config
}

func New(dbList model.Database, cfg pkgConfig.Config) exampleNonDiRepository.Repository {
	return &repository{
		Db:     dbList.Template,
		Config: cfg,
	}
}
