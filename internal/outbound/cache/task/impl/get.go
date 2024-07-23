package impl

import (
	"base/internal/outbound/db/model"
	pkgHelper "base/pkg/helper"
	"context"
	"encoding/json"
)

func (r *repository) GetTaskByID(ctx context.Context, id int64) (task model.Task, err error) {
	span, ctx := pkgHelper.UpdateCtxSpanRepository(ctx)
	defer span.End()

	res, err := r.Cache.Get(ctx, generateTaskByIDKey(id)).Result()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(res), &task)
	if err != nil {
		return
	}

	return
}
