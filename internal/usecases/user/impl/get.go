package impl

import (
	"base/internal/inbound/model"
	dbModel "base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"errors"
)

func (u *usecase) GetUserByID(ctx context.Context, userId int) (user model.Users, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	var userData dbModel.User
	userData, err = u.User.GetUserByID(ctx, userId)

	if err != nil {
		if err.Error() == "data not found" {
			return model.Users{}, errors.New("user not found")
		}
		return model.Users{}, err
	}
	user = model.Users{
		Username: userData.Username,
		Email:    userData.Email,
		UserId:   userData.ID,
	}

	return
}
