package authentication

import (
	activedirectory "base/internal/outbound/authentication/activeDirectory"
	"fmt"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In
}

func Register(container *dig.Container) error {
	err := container.Provide(activedirectory.New)
	if err != nil {
		return fmt.Errorf("[DI] error provide active directory: %+v", err)
	}

	return nil
}
