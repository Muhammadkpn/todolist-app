package petStore

import (
	"base/internal/usecases/model"
	"context"
)

type (
	UseCase interface {
		FindPetByID(ctx context.Context, id uint64) (model.Pet, error)
	}
)
