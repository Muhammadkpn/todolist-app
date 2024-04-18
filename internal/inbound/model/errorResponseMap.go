package model

import (
	pkgError "base/pkg/constant/error"
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrorMap = map[error]int{
		gorm.ErrRecordNotFound:        http.StatusNotFound,
		pkgError.ErrFailedOnValidator: http.StatusBadRequest,
	}
)
