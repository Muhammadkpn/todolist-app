package shared

import "context"

type (
	UseCase interface {
		SomeFunction(ctx context.Context) string
	}
)
