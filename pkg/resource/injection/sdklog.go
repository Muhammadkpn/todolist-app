package injection

import (
	sdklog "gitlab.banksinarmas.com/go/sdkv2/log"
)

func InitSdkLog() sdklog.Logger {
	// init logger
	sdklogCfg := sdklog.Config{
		EnableConsoleLog: true,
		Level:            sdklog.DebugLevel,
	}
	sdklog := sdklog.InitLogger(sdklogCfg)

	return sdklog
}
