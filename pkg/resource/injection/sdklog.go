package injection

import (
	sdkZap "gitlab.banksinarmas.com/go/sdkv2/log/integrations/zap"
	sdkLog "gitlab.banksinarmas.com/go/sdkv2/log/logger"
)

func NewLogger() sdkLog.Logger {
	return sdkZap.New(
		sdkZap.WithLogPath("./logs"),
		sdkZap.WithEnableConsoleLog(true),
		sdkZap.WithLevel(sdkZap.DebugLevel), // dev only, log from debug level
	)
}
