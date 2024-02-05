package impl

import (
	"context"

	"go.elastic.co/apm"
)

func (u *usecase) DeleteTask(ctx context.Context, id int) (err error) {
	span, _ := apm.StartSpan(ctx, "usecase", "DeleteTask")
	defer span.End()

	err = u.TaskRepository.DeleteTask(id)

	return
}
