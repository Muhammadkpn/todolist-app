package petStoreV1

import (
	"base/internal/usecases/model"
	"base/pkg/enums/petType"
	"context"
)

func (*usecase) FindPetByID(ctx context.Context, id uint64) (model.Pet, error) {
	return model.Pet{
		ID:   1,
		Name: "Max",
		Type: petType.Dog,
	}, nil
}
