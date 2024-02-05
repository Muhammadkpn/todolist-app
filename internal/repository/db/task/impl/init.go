package impl

import (
	taskRepository "base/internal/repository/db/task"
	pkgConfig "base/pkg/config"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type repository struct {
	Db     *sqlx.DB
	DbGorm *gorm.DB
	Config pkgConfig.Config
}

func New(db *sqlx.DB, dbGorm *gorm.DB, cfg pkgConfig.Config) taskRepository.Repository {
	return &repository{
		Db:     db,
		DbGorm: dbGorm,
		Config: cfg,
	}
}
