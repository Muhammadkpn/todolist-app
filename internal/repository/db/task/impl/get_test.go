package impl

// import (
// 	"base/internal/model"
// 	"context"
// 	"errors"
// 	"reflect"
// 	"testing"

// 	"github.com/jmoiron/sqlx"
// 	postgres "go.elastic.co/apm/module/apmgormv2/driver/postgres"
// 	"gopkg.in/DATA-DOG/go-sqlmock.v1"
// 	"gorm.io/gorm"
// )

// func Test_repository_GetAllTasks(t *testing.T) {
// 	db, mock, _ := sqlmock.New()
// 	defer db.Close()

// 	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

// 	tests := []struct {
// 		name      string
// 		mock      func()
// 		cfg       model.Config
// 		wantTasks []model.Task
// 		wantErr   bool
// 	}{
// 		{
// 			name: "test success sqlx",
// 			mock: func() {
// 				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow("1", "Task 1"))
// 			},
// 			cfg: model.Config{
// 				FeatureFlag: model.FeatureFlag{
// 					UseGorm: false,
// 				},
// 			},
// 			wantTasks: []model.Task{
// 				{
// 					ID:    1,
// 					Title: "Task 1",
// 				},
// 			},
// 		},
// 		{
// 			name: "test fail sqlx",
// 			mock: func() {
// 				mock.ExpectQuery("SELECT").WillReturnError(errors.New("err"))
// 			},
// 			cfg: model.Config{
// 				FeatureFlag: model.FeatureFlag{
// 					UseGorm: false,
// 				},
// 			},
// 			wantTasks: nil,
// 			wantErr:   true,
// 		},
// 		{
// 			name: "test success gorm",
// 			mock: func() {
// 				mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow("1", "Task 1"))
// 			},
// 			cfg: model.Config{
// 				FeatureFlag: model.FeatureFlag{
// 					UseGorm: true,
// 				},
// 			},
// 			wantTasks: []model.Task{
// 				{
// 					ID:    1,
// 					Title: "Task 1",
// 				},
// 			},
// 		},
// 		{
// 			name: "test fail gorm",
// 			mock: func() {
// 				mock.ExpectQuery("SELECT").WillReturnError(errors.New("err"))
// 			},
// 			cfg: model.Config{
// 				FeatureFlag: model.FeatureFlag{
// 					UseGorm: true,
// 				},
// 			},
// 			wantTasks: nil,
// 			wantErr:   true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mock()
// 			r := &repository{
// 				Db:     sqlx.NewDb(db, "sqlmock"),
// 				DbGorm: gormDB,
// 				Config: tt.cfg,
// 			}
// 			gotTasks, err := r.GetAllTasks(context.Background())
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("repository.GetAllTasks() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotTasks, tt.wantTasks) {
// 				t.Errorf("repository.GetAllTasks() = %v, want %v", gotTasks, tt.wantTasks)
// 			}
// 		})
// 	}
// }
