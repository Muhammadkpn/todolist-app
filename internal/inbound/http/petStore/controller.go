package petStore

import (
	usecase "base/internal/usecase"
	"base/pkg/resource"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	UseCase  usecase.Usecase
	Resource resource.Resource
}

var _ StrictServerInterface = (*Controller)(nil)
