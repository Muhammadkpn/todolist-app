package petStore

import (
	usecase "base/internal/usecase"
	"base/pkg/resource"
	"context"

	"go.uber.org/dig"
)

var (
	tag = `[PetStoreController]`
)

type Controller struct {
	dig.In

	UseCase  usecase.Usecase
	Resource resource.Resource
}

// GetStores implements StrictServerInterface.
func (c *Controller) GetStores(ctx context.Context, request GetStoresRequestObject) (GetStoresResponseObject, error) {
	panic("unimplemented")
}

// PostStores implements StrictServerInterface.
func (c *Controller) PostStores(ctx context.Context, request PostStoresRequestObject) (PostStoresResponseObject, error) {
	panic("unimplemented")
}

var _ StrictServerInterface = (*Controller)(nil)
