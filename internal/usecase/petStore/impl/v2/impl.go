package petStoreV2

import "base/internal/usecase/petStore"

type (
	usecase struct{}
)

func New() petStore.UseCase {
	return &usecase{}
}
