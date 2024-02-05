package injection

import (
	pkgConfig "base/pkg/config"

	"gorm.io/gorm"

	postgres "go.elastic.co/apm/module/apmgormv2/driver/postgres"
)

func InitGorm(cfg pkgConfig.Config) (db *gorm.DB, err error) {
	connectionString := cfg.Database.GenerateConnectionString()
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	return
}
