package pkgDi

import (
	"sync"

	"base/internal/datalogic"
	"base/internal/outbound"
	usecase "base/internal/usecase"
	pkgConfig "base/pkg/config"
	pkgResource "base/pkg/resource"

	"go.uber.org/dig"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	var outer error

	once.Do(func() {
		container = dig.New()

		if err := pkgConfig.Register(container); err != nil {
			outer = err
			return
		}

		if err := pkgResource.Register(container); err != nil {
			outer = err
			return
		}

		if err := usecase.Register(container); err != nil {
			outer = err
			return
		}

		if err := outbound.Register(container); err != nil {
			outer = err

			return
		}

		if err := datalogic.Register(container); err != nil {
			outer = err

			return
		}
	})

	if outer != nil {
		return nil, outer
	}

	return container, nil
}
