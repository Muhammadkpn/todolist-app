package petStore

import (
	"context"
)

func (c *Controller) DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error) {
	return DeletePet204Response{}, nil
}
