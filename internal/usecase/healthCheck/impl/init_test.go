package impl

import (
	pkgConfig "base/pkg/config"
	"base/pkg/resource/model"
	"testing"
)

func TestNew(t *testing.T) {
	New(model.Database{}, model.Redis{}, pkgConfig.Config{})
}
