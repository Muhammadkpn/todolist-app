package pkgError

import "errors"

var (
	ErrUnknownPetType = errors.New("unknown pet type")
)
