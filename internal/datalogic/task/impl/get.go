package impl

import (
	"base/internal/datalogic/model"
	pkgHelper "base/pkg/helper"
	"context"

	sdkZap "gitlab.banksinarmas.com/go/sdkv2/log/integrations/zap"
)

// GetTaskByID implements task.Datalogic.
func (d *datalogic) GetTaskByID(ctx context.Context, id int64) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanDatalogic(ctx)
	defer span.End()

	// get from cache
	taskRepo, err := d.TaskCache.GetTaskByID(ctx, id)
	if err == nil {
		task.MapFromOutbound(taskRepo)

		return
	}

	// only get db if cache not exist
	taskRepo, err = d.TaskDb.GetTaskByID(ctx, id)
	if err != nil {
		return
	}

	// set cache data into db
	err = d.TaskCache.SetTaskByID(ctx, taskRepo)
	if err != nil {
		sdkZap.Error(ctx, "[GetTaskByID][SetTaskByID]Error set cache with data: %+v, error: %+v", taskRepo, err)
	}

	task.MapFromOutbound(taskRepo)

	return
}
