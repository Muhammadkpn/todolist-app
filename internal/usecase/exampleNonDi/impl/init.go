package impl

import (
	exampleNonDiRepository "base/internal/outbound/db/exampleNonDi"
	exampleNonDiUsecase "base/internal/usecase/exampleNonDi"
)

type usecase struct {
	exampleNonDiRepository exampleNonDiRepository.Repository
}

func New(exampleNonDiRepository exampleNonDiRepository.Repository) exampleNonDiUsecase.Usecase {
	return &usecase{
		exampleNonDiRepository: exampleNonDiRepository,
	}
}
