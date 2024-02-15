package user

import (
	usecase "base/internal/usecases"

	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	UseCase usecase.Usecase
}
