package resource

import (
	"base/pkg/resource/injection"
	"fmt"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(injection.NewWebClientFactory); err != nil {
		return fmt.Errorf("[DI] cannot initialize Web Client Factory: %+v", err)
	}

	if err := container.Provide(injection.NewValidator); err != nil {
		return fmt.Errorf("[DI] cannot initialize Validator: %+v", err)
	}

	if err := container.Provide(injection.NewDatabase); err != nil {
		return fmt.Errorf("[DI] cannot initialize Database: %+v", err)
	}

	if err := container.Provide(injection.NewLogger); err != nil {
		return fmt.Errorf("[DI] cannot initialize Logger: %+v", err)
	}

	return nil
}
