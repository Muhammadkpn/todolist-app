package impl

import (
	"base/internal/inbound/model"
	dbModel "base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

func (u *usecase) GetLabel(ctx context.Context, user_id int) (label []model.Labels, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	var dataLabel []dbModel.Label
	if user_id == 0 {
		dataLabel, err = u.Label.GetAllLabel(ctx)
	} else {
		dataLabel, err = u.Label.GetLabelUserId(ctx, user_id)
	}

	for _, dl := range dataLabel {
		lbl := model.Labels{
			CreatedBy:   &dl.CreatedBy,
			CreatedDate: &dl.CreatedDate,
			Id:          &dl.ID,
			Name:        &dl.Name,
			UpdatedBy:   &dl.UpdatedBy,
			UpdatedDate: &dl.UpdatedDate,
			UserId:      &dl.UserID,
		}
		label = append(label, lbl)
	}

	return
}
