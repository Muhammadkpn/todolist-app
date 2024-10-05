package resource

import (
	pkgConfig "base/pkg/config"

	sdkLog "gitlab.banksinarmas.com/go/sdkv2/log/logger"
	sdkValidator "gitlab.banksinarmas.com/go/sdkv2/validator/validator"
	sdkWebClient "gitlab.banksinarmas.com/go/sdkv2/webClient/webClient"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type (
	Resource struct {
		dig.In

		// Configs
		Config pkgConfig.Config

		// WebClientFactory is web client factory to produce WebClient
		WebClientFactory sdkWebClient.Factory

		Logger sdkLog.Logger

		Validator sdkValidator.Validator

		DatabaseSQL *gorm.DB
	}
)
