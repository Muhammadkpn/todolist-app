package impl

import (
	"base/internal/usecases/model"
	"context"

	"go.elastic.co/apm"
)

// import (
// 	"base/internal/model"
// 	"context"

// 	"go.elastic.co/apm"
// )

func (u *usecase) UpdateTask(ctx context.Context, id int, title string) (task model.Task, err error) {
	span, _ := apm.StartSpan(ctx, "usecase", "UpdateTask")
	defer span.End()

	_, err = u.TaskRepository.UpdateTask(id, title)
	if err != nil {
		return model.Task{}, err
	}

	return
}
