package petStoreV2

import "base/internal/usecases/petStore"

type (
	usecase struct{}
)

func New() petStore.UseCase {
	return &usecase{}
}
