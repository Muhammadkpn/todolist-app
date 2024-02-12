package impl

import (
	pkgConfig "base/pkg/config"
	"testing"

	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	New(&gorm.DB{}, pkgConfig.Config{})
}
