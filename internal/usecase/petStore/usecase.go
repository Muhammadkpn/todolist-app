package petStore

import (
	"base/internal/usecase/model"
	"context"
)

type (
	UseCase interface {
		FindPetByID(ctx context.Context, id uint64) (model.Pet, error)
		Create(ctx context.Context, data model.CreatePetRequest) (model.Pet, error)
	}
)
