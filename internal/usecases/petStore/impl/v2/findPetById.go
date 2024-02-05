package petStoreV2

import (
	"base/internal/usecases/model"
	"base/pkg/enums/petType"
	"context"
)

func (*usecase) FindPetByID(ctx context.Context, id uint64) (model.Pet, error) {
	return model.Pet{
		ID:   2,
		Name: "Lily",
		Type: petType.Cat,
	}, nil
}
