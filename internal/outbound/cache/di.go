package cache

import (
	taskCache "base/internal/outbound/cache/task"
	taskCacheImpl "base/internal/outbound/cache/task/impl"
	"fmt"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	TaskCache taskCache.Repository
}

func Register(container *dig.Container) error {
	err := container.Provide(taskCacheImpl.New)
	if err != nil {
		return fmt.Errorf("[DI] error provide task cache: %+v", err)
	}

	return nil
}
