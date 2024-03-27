package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
	"errors"

	"gorm.io/gorm"
)

func (r *repository) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	db := r.DbGorm.WithContext(ctx)
	result := db.Table("task").Where("status = ?", 1).Find(&tasks)
	if result.Error != nil {
		err = result.Error

		return
	}

	return
}

func (r *repository) GetTaskByID(ctx context.Context, id int64) (task model.Task, err error) {
	db := r.DbGorm.WithContext(ctx)
	// task, found := db.Table("task").Get(&task, constant.GetTaskByIDQuery, id)

	err = db.Table("task").Where("id = ?", id).First(&task).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Task{}, errors.New("Data not Found")
		}
		return model.Task{}, err
	}

	return
}
