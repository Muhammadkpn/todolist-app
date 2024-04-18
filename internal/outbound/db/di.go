package db

import (
	"base/internal/outbound/db/petStore"
	taskRepo "base/internal/outbound/db/task"
	taskRepoImpl "base/internal/outbound/db/task/impl"
	"fmt"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	TaskRepository taskRepo.Repository
	PetStore       petStore.Repository
}

func Register(container *dig.Container) error {
	err := container.Provide(taskRepoImpl.New)
	if err != nil {
		return fmt.Errorf("[DI] error provide task repository: %+v", err)
	}

	if err := container.Provide(petStore.New); err != nil {
		return fmt.Errorf("[DI] error provide petStore repository: %+v", err)
	}

	return nil
}
