package task

import (
	modelInbound "base/internal/inbound/model"
	modelUsecase "base/internal/usecases/model"
	"base/internal/usecases/task"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestController_GetTasks(t *testing.T) {
	mockTaskUsecase := task.NewMockUsecase(t)
	tempID := int64(1)
	tempTitle := "aaa"
	tempStatus := 1
	tempErrorCode := ""
	tempErrorMsg := "err"
	tempTask := []modelInbound.Task{
		{
			Id:     &tempID,
			Title:  &tempTitle,
			Status: &tempStatus,
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
				ErrorCode:    &tempErrorCode,
				ErrorMessage: &tempErrorMsg,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			c := &Controller{
				Task: mockTaskUsecase,
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
