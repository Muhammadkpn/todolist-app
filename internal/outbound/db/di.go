package db

import (
	taskRepo "base/internal/outbound/db/task"
	taskRepoImpl "base/internal/outbound/db/task/impl"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	TaskRepository taskRepo.Repository
}

func Register(container *dig.Container) error {
	err := container.Provide(taskRepoImpl.New)
	if err != nil {
		return err
	}

	return nil
}
