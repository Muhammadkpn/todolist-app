package db

import (
	userRepo "base/internal/repository/db/user"
	userImpl "base/internal/repository/db/user/impl"

	taskRepo "base/internal/repository/db/task"
	taskImpl "base/internal/repository/db/task/impl"

	labelRepo "base/internal/repository/db/label"
	labelImpl "base/internal/repository/db/label/impl"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	UserRepository  userRepo.Repository
	taskRepository  taskRepo.Repository
	labelRepository labelRepo.Repository
}

func Register(container *dig.Container) error {
	var err error

	err = container.Provide(userImpl.New)
	if err != nil {
		return err
	}

	err = container.Provide(taskImpl.New)
	if err != nil {
		return err
	}

	err = container.Provide(labelImpl.New)
	if err != nil {
		return err
	}

	return nil
}
