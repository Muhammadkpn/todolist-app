package injection

import (
	sdkHeimdall "gitlab.banksinarmas.com/go/sdkv2/webClient/integration/heimdall"
	sdkWebClient "gitlab.banksinarmas.com/go/sdkv2/webClient/webClient"
)

func NewWebClientFactory() sdkWebClient.Factory {
	return sdkHeimdall.NewClientFactory()
}
