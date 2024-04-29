package model

import (
	pkgConfig "base/pkg/config"
	"base/pkg/constant"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func Test_generateHealthCheckResponse(t *testing.T) {
	type args struct {
		db  pkgConfig.DatabaseCfg
		err error
	}

	cfg := pkgConfig.DatabaseCfg{
		Name:     "template",
		Driver:   "postgres",
		Host:     "localhost",
		Port:     5432,
		Username: "admin",
		Password: "admin",
		Schema:   "template",
	}

	tests := []struct {
		name string
		args args
		want HealthCheckResponse
	}{
		{
			name: "test up",
			args: args{
				db:  cfg,
				err: nil,
			},
			want: HealthCheckResponse{
				Name:   cfg.Name,
				Type:   fmt.Sprintf("%s/%s", constant.ServiceDatabase, cfg.Driver),
				Host:   cfg.Host,
				Status: constant.Up,
			},
		},
		{
			name: "test down",
			args: args{
				db:  cfg,
				err: errors.New("err"),
			},
			want: HealthCheckResponse{
				Name:   cfg.Name,
				Type:   fmt.Sprintf("%s/%s", constant.ServiceDatabase, cfg.Driver),
				Host:   cfg.Host,
				Status: constant.Down,
				Error:  errors.New("err"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateHealthCheckResponse(tt.args.db, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateHealthCheckResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
