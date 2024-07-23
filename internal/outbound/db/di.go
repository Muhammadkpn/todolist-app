package db

import (
	"base/internal/outbound/db/petStore"
	taskDb "base/internal/outbound/db/task"
	taskDbImpl "base/internal/outbound/db/task/impl"
	"fmt"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	TaskDb   taskDb.Repository
	PetStore petStore.Repository
}

func Register(container *dig.Container) error {
	err := container.Provide(taskDbImpl.New)
	if err != nil {
		return fmt.Errorf("[DI] error provide task repository: %+v", err)
	}

	if err := container.Provide(petStore.New); err != nil {
		return fmt.Errorf("[DI] error provide petStore repository: %+v", err)
	}

	return nil
}
