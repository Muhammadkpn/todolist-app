package resource

import (
	pkgConfig "base/pkg/config"

	sdklog "github.com/fericosali/sdk/log"
	sdkWebClient "github.com/fericosali/sdk/webClient/webClient"
	"go.uber.org/dig"
)

type (
	Resource struct {
		dig.In

		// Configs
		Config pkgConfig.Config

		// WebClientFactory is web client factory to produce WebClient
		WebClientFactory sdkWebClient.Factory

		Sdklog sdklog.Logger
	}
)
