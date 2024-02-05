package pkgConfig

import (
	"fmt"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewConfig); err != nil {
		return fmt.Errorf("[DI] cannot initialize config: %+v", err)
	}

	return nil
}
