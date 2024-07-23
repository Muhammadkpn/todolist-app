package outbound

import (
	cacheRepo "base/internal/outbound/cache"
	dbRepo "base/internal/outbound/db"

	"go.uber.org/dig"
)

type Outbound struct {
	dig.In

	Repository      dbRepo.Repository
	CacheRepository cacheRepo.Repository
}

func Register(container *dig.Container) (err error) {
	err = dbRepo.Register(container)
	if err != nil {
		return err
	}

	err = cacheRepo.Register(container)
	if err != nil {
		return
	}

	return nil
}
