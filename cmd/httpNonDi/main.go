package main

import (
	sdkRunner "gitlab.banksinarmas.com/go/sdkv2/appRunner"
	sdkServer "gitlab.banksinarmas.com/go/sdkv2/appRunner/server"

	innerRoute "base/internal/inbound/http"
)

// The main function serves as the entry point for the application.
// It initializes and runs the server using the Run function and handles any errors.
func main() {
	InitModule()

	appCfg := Config.AppConfig
	sdkRunner.New().
		With(sdkServer.NewHttp(appCfg.ServiceName, appCfg.ServiceVersion, appCfg.BasePath,
			sdkServer.WithHttpPort(appCfg.ServerPort),
			sdkServer.WithHttpGraceFulPeriod(appCfg.GracefullyDuration),
			sdkServer.WithHttpRegister(innerRoute.Routes),
		)).
		Run()
}
