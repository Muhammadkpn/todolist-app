package http

import (
	"base/pkg/constant"

	"github.com/labstack/echo/v4"
	sdkZap "gitlab.banksinarmas.com/go/sdkv2/log/integrations/zap"
)

// Handle the data here (e.g., print to console, write to file)
func myBodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	sdkZap.Debug(c.Request().Context(), "[Middleware][BodyDump]Endpoint: %+v, Method: %+v, Request: %+v, Response: %+v", c.Path(), c.Request().Method, string(reqBody), string(resBody))
}

// Define a skipper function to exclude specific routes (optional)
func skipBodyDump(c echo.Context) bool {
	return !constant.MapWhitelistLogEndpoint[c.Path()+"_"+c.Request().Method]
}
