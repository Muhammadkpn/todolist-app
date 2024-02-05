package petStore

import (
	"context"
)

func (c *Controller) AddPet(ctx context.Context, request AddPetRequestObject) (AddPetResponseObject, error) {
	data, err := c.UseCase.PetStoreV1.FindPetByID(ctx, 1)
	if err != nil {
		return AddPetdefaultJSONResponse{}, nil
	}

	return AddPet200JSONResponse{
		Id:      int64(data.ID),
		Name:    data.Name,
		PetType: &data.Type,
	}, nil
}
