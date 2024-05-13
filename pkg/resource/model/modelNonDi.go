package model

import (
	exampleNonDiController "base/internal/inbound/http/exampleNonDi"
)

type Controller struct {
	ControllerA *exampleNonDiController.Controller
}

var ControllerList Controller
