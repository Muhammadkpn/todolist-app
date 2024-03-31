package petStore

import (
	usecase "base/internal/usecase"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	UseCase usecase.Usecase
}

var _ StrictServerInterface = (*Controller)(nil)
