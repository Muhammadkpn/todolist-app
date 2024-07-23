package model

import (
	obModel "base/internal/outbound/db/model"
	"base/pkg/enums/petType"
)

type (
	Pet struct {
		ID      uint64
		Name    string
		PetType petType.Type
	}

	CreatePetRequest struct {
		Name    string       `validate:"required"`
		PetType petType.Type `validate:"required"`
	}
)

func (p Pet) FromDBModel(data obModel.Pet) Pet {
	return Pet{
		ID:      data.ID,
		Name:    data.Name,
		PetType: data.PetType,
	}
}
