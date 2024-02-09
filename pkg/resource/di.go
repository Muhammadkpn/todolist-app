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

	if err := container.Provide(injection.InitGorm); err != nil {
		return fmt.Errorf("[DI] cannot initialize Gorm: %+v", err)
	}

	if err := container.Provide(injection.InitSdkLog); err != nil {
		return fmt.Errorf("[DI] cannot initialize Sdklog: %+v", err)
	}

	return nil
}
