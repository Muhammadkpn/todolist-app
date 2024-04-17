package task

import (
	modelInbound "base/internal/inbound/model"
	modelUsecase "base/internal/usecase/model"
	"base/internal/usecase/task"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestController_GetTasks(t *testing.T) {
	mockTaskUsecase := task.NewMockUsecase(t)
	tempTask := []modelInbound.Task{
		{
			Id:     1,
			Title:  "aaa",
			Status: 1,
		},
	}

	type args struct {
		ctx     context.Context
		request GetTasksRequestObject
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    GetTasksResponseObject
		wantErr bool
	}{
		{
			name: "test success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockTaskUsecase.On("GetAllTasks", mock.Anything).Return([]modelUsecase.Task{
					{
						ID:     1,
						Title:  "aaa",
						Status: 1,
					},
				}, nil).Once()
			},
			want: GetTasks200JSONResponse{
				Data: &tempTask,
			},
			wantErr: false,
		},
		{
			name: "test error",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockTaskUsecase.On("GetAllTasks", mock.Anything).Return([]modelUsecase.Task{}, errors.New("err")).Once()
			},
			want: GetTasks500JSONResponse{
				Message: "err",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			c := &Controller{
				TaskUsecase: mockTaskUsecase,
			}
			got, err := c.GetTasks(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.GetTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
