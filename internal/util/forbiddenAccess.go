package util

import (
	"errors"
	"strconv"
	"strings"
)

func ForbiddenAccess(access string, role int) (err error) {
	roles := "[" + strconv.Itoa(role) + "]"

	if !strings.Contains(access, roles) {
		err = errors.New("you are not allowed to access this resource")
	}
	return
}
