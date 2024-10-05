package impl

import (
	//	"base/internal/usecases/model"
	"base/internal/usecases/model"
	"context"
)

func (u *usecase) GetClaims(ctx context.Context, reqToken string) (claim model.DataJwt, err error) {
	claims, _, err := u.CheckToken(ctx, reqToken)
	if err != nil {
		return
	}

	claim.Id = claims["Id"].(string)
	claim.Username = claims["Username"].(string)
	claim.Email = claims["Email"].(string)

	return
}
