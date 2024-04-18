package pkgHelper

import (
	pkgError "base/pkg/constant/error"
	"errors"
	"net/http"

	sdkError "gitlab.banksinarmas.com/go/sdkv2/common/error"
)

func FromErrorMap(err error, errMap map[error]int) int {
	var err2 error = err
	if _, ok := err.(sdkError.Err); ok {
		err2 = err.(sdkError.Err).RootError()
	}

	if errors.Is(err2, pkgError.ErrFailedOnValidator) {
		return http.StatusBadRequest
	}

	if f, ok := errMap[err2]; ok {
		return f
	}

	return http.StatusInternalServerError
}
