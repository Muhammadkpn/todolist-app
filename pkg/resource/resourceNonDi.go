package resource

import (
	pkgConfig "base/pkg/config"
	"base/pkg/resource/injection"
	"base/pkg/resource/model"
)

type (
	ResourceNonDi struct {
		Config      pkgConfig.Config
		DatabaseSQL model.Database
	}
)

func InitResourceNonDi() ResourceNonDi {
	cfg, err := pkgConfig.NewConfig()
	if err != nil {
		panic("Error on init config")
	}

	db, err := injection.NewDatabase(cfg)
	if err != nil {
		panic("Error on init database")
	}

	return ResourceNonDi{
		Config:      cfg,
		DatabaseSQL: db,
	}
}
