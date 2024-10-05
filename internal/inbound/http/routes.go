package http

import (
	"base/internal/inbound/http/labels"
	"base/internal/inbound/http/tasks"
	"base/internal/inbound/http/users"
	pkgResource "base/pkg/resource"

	sdkHttpMiddleware "gitlab.banksinarmas.com/go/sdkv2/appRunner/middleware/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Resource pkgResource.Resource

	Users  users.Controller
	Labels labels.Controller
	Tasks  tasks.Controller
}

func (i Inbound) Routes(ec *echo.Echo) {
	ec.Use(sdkHttpMiddleware.AppContext())

	// middleware for log req and res
	ec.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: skipBodyDump,
		Handler: myBodyDumpHandler,
	}))

	ec.Validator = i.Resource.Validator
	base := ec.Group(i.Resource.Config.AppConfig.BasePath)

	v1 := base.Group("/v1")
	users.RegisterHandlers(v1, users.NewStrictHandler(&i.Users, nil))
	labels.RegisterHandlers(v1, labels.NewStrictHandler(&i.Labels, nil))
	tasks.RegisterHandlers(v1, tasks.NewStrictHandler(&i.Tasks, nil))
}
