package healthCheck

import (
	"go.uber.org/dig"

	healthCheckUsecase "base/internal/usecase/healthCheck"
)

type Controller struct {
	dig.In
	HealthCheckUsecase healthCheckUsecase.Usecase
}
