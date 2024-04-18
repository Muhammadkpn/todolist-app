package main

import (
	"base/internal/inbound"
	pkgDi "base/pkg/di"
	pkgResource "base/pkg/resource"
	"context"

	sdkRunner "gitlab.banksinarmas.com/go/sdkv2/appRunner"
	sdkServer "gitlab.banksinarmas.com/go/sdkv2/appRunner/server"
	sdkLogger "gitlab.banksinarmas.com/go/sdkv2/log/logger"
	"go.uber.org/dig"
)

type (
	ContainerCall func() (*dig.Container, error)
	Invoke        func(container *dig.Container) error
	InvokeError   func(container *dig.Container, err error)
)

// Run initializes a DI container, invokes a function with dependencies,
// and handles errors by calling the provided onError function.
func Run(containerCall ContainerCall, invoke Invoke, onError InvokeError) error {
	container, err := containerCall()
	if err != nil {
		return err
	}

	if err := invoke(container); err != nil {
		onError(container, err)
	}

	return nil
}

func run(container *dig.Container) error {
	return container.Invoke(func(inbound inbound.Inbound, resource pkgResource.Resource) error {
		// Initializes and starts an HTTP server with the specified port
		appCfg := resource.Config.AppConfig
		sdkRunner.New().
			With(sdkServer.NewHttp(appCfg.ServiceName, appCfg.ServiceVersion, appCfg.BasePath,
				sdkServer.WithHttpPort(appCfg.ServerPort),
				sdkServer.WithHttpGraceFulPeriod(appCfg.GracefullyDuration),
				sdkServer.WithHttpRegister(inbound.Http.Routes),
			)).
			Run()

		return nil
	})
}

// onError is a callback function called when an error occurs during
// the execution of the main function. It attempts to invoke a function
// for additional error handling and prints the error if it persists.
func onError(container *dig.Container, err error) {
	err = container.Invoke(func(logger sdkLogger.Logger) error {
		if err != nil {
			logger.Info(context.Background(), err.Error())
			return err
		}
		return nil
	})
}

// The main function serves as the entry point for the application.
// It initializes and runs the server using the Run function and handles any errors.
func main() {
	// Example usage of the Run function to start the server.
	if err := Run(pkgDi.Container, run, onError); err != nil {
		panic(err)
	}
}
