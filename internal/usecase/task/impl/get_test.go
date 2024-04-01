package impl

import (
	modelDB "base/internal/repository/db/model"
	"base/internal/repository/db/task"
	"base/internal/usecase/model"
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_usecase_GetAllTasks(t *testing.T) {
	mockdb := task.NewMockRepository(t)

	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name      string
		mock      func()
		wantTasks []model.Task
		wantErr   bool
		args      args
	}{
		{
			name: "test success",
			mock: func() {
				mockdb.On("GetAllTasks", mock.Anything).Return([]modelDB.Task{
					{
						ID:     1,
						Title:  "test",
						Status: 1,
					},
				}, nil).Once()
			},
			wantTasks: []model.Task{
				{
					ID:     1,
					Title:  "test",
					Status: 1,
				},
			},
			args: args{
				ctx: context.Background(),
			},
		},
		{
			name: "test fail",
			mock: func() {
				mockdb.On("GetAllTasks", mock.Anything).Return([]modelDB.Task{}, errors.New("err")).Once()
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			u := &usecase{
				TaskRepository: mockdb,
			}
			gotTasks, err := u.GetAllTasks(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("usecase.GetAllTasks() = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}
