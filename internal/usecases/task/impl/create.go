package impl

import (
	dbModel "base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"time"
)

// CreateTask(ctx context.Context, request model.Tasks) (err error)
func (u *usecase) CreateTask(ctx context.Context, user_id int, username string, title string, description string, status string, due_date time.Time, labels []int) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	var label []dbModel.Label
	for _, label_id := range labels {
		lbl := dbModel.Label{
			ID: label_id,
		}
		label = append(label, lbl)
	}

	task := dbModel.Task{
		Title:       title,
		Description: description,
		UserID:      user_id,
		Label:       label,
		Status:      status,
		DueDate:     due_date,
		CreatedBy:   username,
		CreatedDate: time.Now(),
		UpdatedBy:   username,
		UpdatedDate: time.Now(),
	}

	// CreateTask(ctx context.Context, request model.Task) (task model.Task, err error)
	err = u.Task.CreateTask(ctx, task)
	if err != nil {
		return err
	}
	return
}
