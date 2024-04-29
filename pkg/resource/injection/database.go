package injection

import (
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"

	"gorm.io/gorm"

	oracle "github.com/godoes/gorm-oracle"
	postgres "go.elastic.co/apm/module/apmgormv2/driver/postgres"
)

// example with postgres and oracle, please comment unused database
func NewDatabase(cfg pkgConfig.Config) (db model.Database, err error) {
	template, err := gorm.Open(postgres.Open(cfg.Database.Template.GenerateConnectionString()), &gorm.Config{})
	if err != nil {
		return
	}

	oracle, err := gorm.Open(oracle.Open(cfg.Database.Oracle.GenerateConnectionString()), &gorm.Config{})
	if err != nil {
		return
	}

	db = model.Database{
		Template: template,
		Oracle:   oracle,
	}

	return
}
