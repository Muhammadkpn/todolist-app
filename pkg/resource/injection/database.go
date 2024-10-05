package injection

import (
	"base/internal/repository/db/model"
	pkgConfig "base/pkg/config"

	"gorm.io/gorm"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
)

func NewDatabase(cfg pkgConfig.Config) (db *gorm.DB, err error) {
	connectionString := cfg.Database.GenerateConnectionString()
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return
	}

	db = db.Debug()
	// sqlDB, err := db.DB()

	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetMaxIdleConns(10)
	//Use it when to want to run service the first time
	db.AutoMigrate(&model.Label{}, &model.Task{}, &model.TaskLabel{}, &model.User{})

	return

}
