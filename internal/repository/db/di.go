package db

import (
	taskRepo "base/internal/repository/db/task"
	taskRepoImpl "base/internal/repository/db/task/impl"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	TaskRepository taskRepo.Repository
}

func Register(container *dig.Container) error {
	var err error

	err = container.Provide(taskRepoImpl.New)
	if err != nil {
		return err
	}

	return nil
}
