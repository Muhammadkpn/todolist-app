package impl

import (
	"base/internal/outbound/db/model"
	pkgHelper "base/pkg/helper"
	"context"
	"encoding/json"
	"time"
)

func (r *repository) SetTaskByID(ctx context.Context, task model.Task) (err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	raw, err := json.Marshal(task)
	if err != nil {
		return
	}

	r.Cache.Del(ctx, generateTaskByIDKey(task.ID))

	err = r.Cache.SetEx(ctx, generateTaskByIDKey(task.ID), string(raw), 60*time.Second).Err()
	if err != nil {
		return
	}

	return
}
