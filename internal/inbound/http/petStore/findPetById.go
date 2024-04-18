package petStore

import (
	"context"
)

func (c *Controller) FindPetByID(ctx context.Context, request FindPetByIDRequestObject) (FindPetByIDResponseObject, error) {
	data, err := c.UseCase.PetStoreV1.FindPetByID(ctx, 1)
	if err != nil {
		return FindPetByIDdefaultJSONResponse{}, nil
	}

	return FindPetByID200JSONResponse{
		Id:      int64(data.ID),
		Name:    data.Name,
		PetType: data.PetType,
	}, nil
}
