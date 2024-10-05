package impl

import (
	dbModel "base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"log"
	"time"
)

// CreateTask(ctx context.Context, request model.Tasks) (err error)
func (u *usecase) EditTask(ctx context.Context, user_id int, username string, task_id int, title string, description string, status string, due_date time.Time, labels []int) (err error) {
	span, ctx := util.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	var label []dbModel.Label
	for _, label_id := range labels {
		lbl := dbModel.Label{
			ID: label_id,
		}
		log.Println("lbl: ", lbl)
		label = append(label, lbl)
	}
	log.Println("label arr: ", label)

	task := dbModel.Task{
		ID:          task_id,
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
	err = u.Task.UpdateTask(ctx, task)
	if err != nil {
		return err
	}
	return
}
