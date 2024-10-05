package impl

import (
	"base/internal/inbound/model"
	"base/internal/util"
	"context"
)

func (u *usecase) GetTask(ctx context.Context, filter map[string]interface{}, sortBy string, order string) (tasks []model.Tasks, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	result, err := u.Task.GetTask(ctx, filter, sortBy, order)
	if err != nil {
		return
	}

	for _, data := range result {
		var dataLabel []model.DataLabels
		for _, data2 := range data.Label {
			lbl := model.DataLabels{
				CreatedBy:   &data2.CreatedBy,
				CreatedDate: &data2.CreatedDate,
				Id:          data2.ID,
				Name:        &data2.Name,
				UpdatedBy:   &data2.UpdatedBy,
				UpdatedDate: &data2.UpdatedDate,
				UserId:      &data2.UserID,
			}
			dataLabel = append(dataLabel, lbl)
		}
		item := model.Tasks{
			CreatedBy:   &data.CreatedBy,
			CreatedDate: &data.CreatedDate,
			Id:          data.ID,
			Description: &data.Description,
			Title:       data.Title,
			Status:      data.Status,
			UpdatedBy:   &data.UpdatedBy,
			UpdatedDate: &data.UpdatedDate,
			UserId:      &data.UserID,
			Labels:      &dataLabel,
		}
		tasks = append(tasks, item)
	}

	return
}

func (u *usecase) GetTaskById(ctx context.Context, taskId int) (tasks model.Tasks, err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	result, err := u.Task.GetTaskId(ctx, taskId)

	var dataLabel []model.DataLabels
	for _, data2 := range result.Label {
		lbl := model.DataLabels{
			CreatedBy:   &data2.CreatedBy,
			CreatedDate: &data2.CreatedDate,
			Id:          data2.ID,
			Name:        &data2.Name,
			UpdatedBy:   &data2.UpdatedBy,
			UpdatedDate: &data2.UpdatedDate,
			UserId:      &data2.UserID,
		}
		dataLabel = append(dataLabel, lbl)
	}
	tasks = model.Tasks{
		CreatedBy:   &result.CreatedBy,
		CreatedDate: &result.CreatedDate,
		Id:          result.ID,
		Description: &result.Description,
		Title:       result.Title,
		Status:      result.Status,
		UpdatedBy:   &result.UpdatedBy,
		UpdatedDate: &result.UpdatedDate,
		UserId:      &result.UserID,
		Labels:      &dataLabel,
	}

	return
}
