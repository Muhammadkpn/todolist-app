package model

import "base/pkg/enums/petType"

type (
	Pet struct {
		ID   uint64
		Name string
		Type petType.Type
	}
)
