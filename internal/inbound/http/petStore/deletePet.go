package petStore

import (
	"context"
	"fmt"
)

func (c *Controller) DeletePet(ctx context.Context, request DeletePetRequestObject) (DeletePetResponseObject, error) {
	fmt.Println(c.UseCase.Shared.SomeFunction(ctx))
	return DeletePet204Response{}, nil
}
