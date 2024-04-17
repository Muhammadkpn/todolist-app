package outbound

import (
	dbRepo "base/internal/outbound/db"

	"go.uber.org/dig"
)

type Repository struct {
	dig.In

	Db dbRepo.Repository
}

func Register(container *dig.Container) error {
	err := dbRepo.Register(container)
	if err != nil {
		return err
	}

	return nil
}
