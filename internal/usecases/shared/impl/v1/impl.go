package sharedV1

import "base/internal/usecases/shared"

type (
	usecase struct {
	}
)

func New() shared.UseCase {
	return &usecase{}
}
