package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"strings"

	"gorm.io/gorm"
)

func (r *repository) GetTask(ctx context.Context, filter map[string]interface{}, sortBy string, order string) (tasks []model.Task, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx).Model(&model.Task{})

	// Apply filters
	if status, ok := filter["status"]; ok && status != nil {
		db = db.Where("status = ?", status)
	}

	if labelID, ok := filter["labelId"]; ok && labelID != nil {
		db = db.Joins("JOIN task_label ON task_label.task_id = task.id").
			Where("task_label.label_id = ?", labelID)
	}

	if userID, ok := filter["userId"]; ok && userID != nil {
		db = db.Where("user_id = ?", userID)
	}

	// Apply sorting
	if sortBy != "" {
		sortField := ""
		switch sortBy {
		case "date":
			sortField = "due_date"
		case "title":
			sortField = "title"
		default:
			sortField = "id" // Default sort field
		}

		sortOrder := "ASC"
		if strings.EqualFold("desc", order) {
			sortOrder = "DESC"
		}

		db = db.Order(sortField + " " + sortOrder)
	} else {
		db = db.Order("id ASC") // Default order
	}

	// Execute the query
	result := db.Preload("Label").Find(&tasks)
	err = util.HandleError(result)

	return
}

func (r *repository) GetTaskId(ctx context.Context, id int) (task model.Task, err error) {
	// Creating a span for tracing context (if necessary)
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	// Executing the query to find the task by ID and preload the associated labels
	result := r.DbGorm.WithContext(ctx).Preload("Label").First(&task, id)

	// Handling errors that might occur during the query
	if err = result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Task not found
			return task, err
		}
		// Another error occurred
		return task, err
	}

	// Returning the found task with associated labels
	return task, nil
}
