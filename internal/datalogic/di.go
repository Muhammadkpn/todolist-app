package datalogic

import (
	"go.uber.org/dig"

	taskDatalogicImpl "base/internal/datalogic/task/impl"
)

func Register(container *dig.Container) (err error) {
	err = container.Provide(taskDatalogicImpl.New)
	if err != nil {
		return err
	}

	return nil
}
