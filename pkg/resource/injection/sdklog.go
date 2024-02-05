package injection

import (
	sdklog "github.com/fericosali/sdk/log"
)

func InitSdkLog() sdklog.Logger {
	// init logger
	sdklogCfg := sdklog.Config{
		LogPath:          "./../logs",
		EnableConsoleLog: true,
		Level:            sdklog.DebugLevel,
	}
	sdklog := sdklog.InitLogger(sdklogCfg)

	return sdklog
}
