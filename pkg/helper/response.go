package pkgHelper

import (
	"net/http"
)

func FromErrorMap(err error, errMap map[error]int) int {
	if f, ok := errMap[err]; ok {
		return f
	}

	return http.StatusInternalServerError
}
