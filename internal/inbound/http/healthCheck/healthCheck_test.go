package healthCheck

import (
	util "base/internal/usecase/healthCheck"
	"base/pkg/resource/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func TestController_HealthCheck(t *testing.T) {
	mockUsecase := new(util.MockUsecase)
	type args struct {
		eCtx echo.Context
	}

	// Create an instance of Echo
	e := echo.New()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Create a new echo context
	eCtx := e.NewContext(req, rec)

	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "testt success",
			mock: func() {
				mockUsecase.On("HealthCheck", mock.Anything).Return([]model.HealthCheckResponse{}, nil).Once()
			},
			args: args{
				eCtx: eCtx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			c := &Controller{
				HealthCheckUsecase: mockUsecase,
			}
			if err := c.HealthCheck(tt.args.eCtx); (err != nil) != tt.wantErr {
				t.Errorf("Controller.HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
