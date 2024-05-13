package user

import (
	exampleNonDiUsecase "base/internal/usecase/exampleNonDi"
)

type Controller struct {
	ExampleNonDiUsecase exampleNonDiUsecase.Usecase
}
