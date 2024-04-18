package pkgHelper

import (
	"net/http"
)

func FromErrorMap(msg string, errMap map[string]int) int {
	if f, ok := errMap[msg]; ok {
		return f
	}

	return http.StatusInternalServerError
}
