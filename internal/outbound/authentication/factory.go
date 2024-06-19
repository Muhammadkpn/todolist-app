package authentication

import (
	activedirectoryAuth "base/internal/outbound/authentication/activeDirectory"
	directAuth "base/internal/outbound/authentication/direct"
	interfaces "base/internal/outbound/interface"
	pkgConfig "base/pkg/config"
	"errors"
)

const (
	ActiveDirectory = "active_directory"
	Direct          = "direct"
)

var (
	repositories         = map[string]interfaces.Authentication{}
	ErrInvalidAuthMethod = errors.New("invalid authentication method")
)

func InitRepositories(cfg pkgConfig.Config) map[string]interfaces.Authentication {
	repositories := make(map[string]interfaces.Authentication)
	repositories[ActiveDirectory] = activedirectoryAuth.New(cfg)
	repositories[Direct] = directAuth.New(cfg)

	return repositories
}
