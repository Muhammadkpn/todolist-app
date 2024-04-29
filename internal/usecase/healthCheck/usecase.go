package util

import (
	"base/pkg/resource/model"
	"context"
)

type Usecase interface {
	HealthCheck(ctx context.Context) (res []model.HealthCheckResponse, err error)
}
