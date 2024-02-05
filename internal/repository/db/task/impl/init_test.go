package impl

import (
	pkgConfig "base/pkg/config"
	"testing"

	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	New(&sqlx.DB{}, &gorm.DB{}, pkgConfig.Config{})
}
