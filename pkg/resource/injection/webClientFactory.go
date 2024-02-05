package injection

import (
	sdkHeimdall "github.com/fericosali/sdk/webClient/integration/heimdall"
	sdkWebClient "github.com/fericosali/sdk/webClient/webClient"
)

func NewWebClientFactory() sdkWebClient.Factory {
	return sdkHeimdall.NewClientFactory()
}
