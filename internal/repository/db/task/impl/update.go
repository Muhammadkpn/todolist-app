package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"fmt"
	"time"
)

func (r *repository) UpdateTask(ctx context.Context, request model.Task) (err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)

	req := model.Task{
		Title:       request.Title,
		Description: request.Description,
		Status:      request.Status,
		Label:       request.Label,
		DueDate:     request.DueDate,
		UpdatedBy:   request.UpdatedBy,
		UpdatedDate: time.Now(),
	}

	result := db.Model(&model.Task{}).Omit("Label.*").Where("\"id\" = ?", request.ID).Updates(&req)
	err = util.HandleError(result)

	// Collect label IDs from request
	var labelIDs []int
	for _, label := range request.Label {
		labelIDs = append(labelIDs, label.ID)
	}

	// Check if all label IDs exist in the database
	var count int64
	err = db.Model(&model.Label{}).Where("id IN (?)", labelIDs).Count(&count).Error
	if err != nil {
		return err
	}

	// Compare count with the number of labels in the request
	if int(count) != len(labelIDs) {
		return fmt.Errorf("one or more label IDs are invalid or do not exist")
	}

	// Add new labels and remove old labels in the task_label table
	err = db.Model(&request).Association("Label").Replace(request.Label)
	if err != nil {
		return
	}
	return
}
