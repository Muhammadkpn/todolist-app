package petStore

import (
	"base/internal/outbound/model"
	pkgResource "base/pkg/resource"
	"context"

	"gorm.io/gorm"
)

var (
	tag = `[PetStoreRepository]`
)

type (
	Repository interface {
		Create(ctx context.Context, db *gorm.DB, data model.Pet) (model.Pet, error)
	}

	repository struct {
		resource pkgResource.Resource
	}
)

func New(resource pkgResource.Resource) Repository {
	return &repository{
		resource: resource,
	}
}
