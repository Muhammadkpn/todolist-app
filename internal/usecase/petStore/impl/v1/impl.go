package petStoreV1

import (
	"base/internal/outbound"
	"base/internal/usecase/petStore"
	pkgResource "base/pkg/resource"
)

type (
	usecase struct {
		outbound outbound.Outbound
		resource pkgResource.Resource
	}
)

func New(outbound outbound.Outbound, resource pkgResource.Resource) petStore.UseCase {
	return &usecase{
		outbound: outbound,
		resource: resource,
	}
}
