package impl

import (
	"base/internal/outbound/db/model"
	"context"
	"errors"
	"reflect"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_GetAllTasks(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	tests := []struct {
		name      string
		mock      func()
		wantTasks []model.Task
		wantErr   bool
	}{
		{
			name: "test success gorm",
			mock: func() {
				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow("1", "Task 1"))
			},
			wantTasks: []model.Task{
				{
					ID:    1,
					Title: "Task 1",
				},
			},
		},
		{
			name: "test fail gorm",
			mock: func() {
				mock.ExpectQuery("SELECT").WillReturnError(errors.New("err"))
			},
			wantTasks: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			r := &repository{
				DbGorm: gormDB,
			}
			gotTasks, err := r.GetAllTasks(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
				t.Errorf("repository.GetAllTasks() = %v, want %v", gotTasks, tt.wantTasks)
			}
		})
	}
}
