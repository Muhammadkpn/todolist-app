package sharedV1

import "base/internal/usecase/shared"

type (
	usecase struct {
	}
)

func New() shared.UseCase {
	return &usecase{}
}
