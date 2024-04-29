package impl

import (
	pkgConfig "base/pkg/config"
	"base/pkg/constant"
	"base/pkg/resource/model"
	"context"
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_usecase_HealthCheck(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	cfg := pkgConfig.Config{
		Database: pkgConfig.Database{
			Template: pkgConfig.DatabaseCfg{
				Name:     "template",
				Driver:   "postgres",
				Host:     "localhost",
				Port:     5432,
				Username: "admin",
				Password: "admin",
				Schema:   "template",
			},
			Oracle: pkgConfig.DatabaseCfg{
				Name:     "oracle",
				Driver:   "oracle",
				Host:     "localhost",
				Port:     5432,
				Username: "admin",
				Password: "admin",
				Schema:   "oracle",
			},
		},
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantRes []model.HealthCheckResponse
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {

			},
			wantRes: []model.HealthCheckResponse{
				{
					Name:   cfg.Database.Template.Name,
					Host:   cfg.Database.Template.Host,
					Type:   fmt.Sprintf("%s/%s", constant.ServiceDatabase, cfg.Database.Template.Driver),
					Status: constant.Up,
				},
				{
					Name:   cfg.Database.Oracle.Name,
					Host:   cfg.Database.Oracle.Host,
					Type:   fmt.Sprintf("%s/%s", constant.ServiceDatabase, cfg.Database.Oracle.Driver),
					Status: constant.Up,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			u := &usecase{
				Db: model.Database{
					Template: gormDB,
					Oracle:   gormDB,
				},
				cfg: cfg,
			}
			gotRes, err := u.HealthCheck(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("usecase.HealthCheck() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
