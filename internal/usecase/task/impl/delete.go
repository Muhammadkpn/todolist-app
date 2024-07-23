package impl

import (
	pkgHelper "base/pkg/helper"
	"context"
)

func (u *usecase) DeleteTask(ctx context.Context, id int64) (err error) {
	span, ctx := pkgHelper.UpdateCtxSpanUsecase(ctx)
	defer span.End()

	err = u.TaskDb.DeleteTask(ctx, id)

	return
}
