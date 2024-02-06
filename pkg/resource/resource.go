package resource

import (
	pkgConfig "base/pkg/config"

	sdklog "gitlab.banksinarmas.com/go/sdkv2/log"
	sdkWebClient "gitlab.banksinarmas.com/go/sdkv2/webClient/webClient"
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
