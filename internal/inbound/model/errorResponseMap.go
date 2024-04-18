package model

import (
	pkgError "base/pkg/constant/error"
	"net/http"
)

var (
	ErrorMap = map[string]int{
		pkgError.ERR_FAILED_ON_VALIDATOR: http.StatusBadRequest,
	}
)
