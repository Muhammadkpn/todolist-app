package impl

import (
	"context"

	"go.elastic.co/apm"
)

func (u *usecase) DeleteTask(ctx context.Context, id int64) (err error) {
	span, _ := apm.StartSpan(ctx, "usecase", "DeleteTask")
	defer span.End()

	err = u.TaskRepository.DeleteTask(ctx, id)

	return
}
