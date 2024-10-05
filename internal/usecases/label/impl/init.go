package impl

import (
	"base/internal/repository/db/label"
	ujwt "base/internal/usecases/jwt"
	labelUsecase "base/internal/usecases/label"
)

type usecase struct {
	Label label.Repository
	Ujwt  ujwt.Usecase
}

func New(labelRepository label.Repository, ujwt ujwt.Usecase) labelUsecase.Usecase {
	return &usecase{
		Label: labelRepository,
		Ujwt:  ujwt,
	}
}
