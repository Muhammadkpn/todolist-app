package inbound

import (
	"base/internal/inbound/http"

	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Http http.Inbound
}
