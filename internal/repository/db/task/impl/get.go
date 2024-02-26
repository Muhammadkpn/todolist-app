package impl

import (
	"base/internal/repository/db/model"
	"base/internal/util"
	"context"
)

func (r *repository) GetAllTasks(ctx context.Context) (tasks []model.Task, err error) {
	span, ctx := util.UpdateCtxSpanRepository(ctx)
	defer span.End()

	// time.Sleep(1 * time.Second)

	db := r.DbGorm.WithContext(ctx)
	result := db.Table("task").Where("status = ?", 1).Find(&tasks)
	if result.Error != nil {
		err = result.Error

		return
	}

	// if !r.Config.FeatureFlag.UseGorm {
	// 	// use sqlx
	// 	err = r.Db.Select(&tasks, constant.GetAllTaskQuery)
	// 	if err != nil {
	// 		return
	// 	}
	// } else {
	// 	// use gorm
	// 	db := r.DbGorm.WithContext(ctx)

	// 	time.Sleep(2 * time.Second)
	// 	result := db.Table("task").Where("status = ?", 1).Find(&tasks)
	// 	if result.Error != nil {
	// 		err = result.Error

	// 		return
	// 	}
	// }

	return
}

func (r *repository) GetTaskByID(id int64) (task model.Task, err error) {
	// err = r.Db.Get(&task, constant.GetTaskByIDQuery, id)

	return
}
