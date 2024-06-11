package http

import (
	"base/pkg/constant"
	"fmt"

	"github.com/labstack/echo/v4"
)

func myBodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	// Handle the data here (e.g., print to console, write to file)
	fmt.Println("Request Body:", string(reqBody))
	fmt.Println("Response Body:", string(resBody))
}

// Define a skipper function to exclude specific routes (optional)
func skipBodyDump(c echo.Context) bool {
	return !constant.MapWhitelistLogEndpoint[c.Path()+"_"+c.Request().Method]
}
