package model

import (
	"net/http"

	"gorm.io/gorm"
)

var (
	ErrorMap = map[error]int{
		gorm.ErrRecordNotFound: http.StatusNotFound,
	}
)
