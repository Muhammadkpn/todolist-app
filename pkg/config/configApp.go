package pkgConfig

import "time"

type (
	AppConfig struct {
		ServiceVersion     string        `yaml:"version" required:"true"`
		ServiceName        string        `yaml:"name" required:"true"`
		BasePath           string        `yaml:"base-path" default:"sample-project"`
		ServerPort         int           `yaml:"port" default:"8090"`
		GracefullyDuration time.Duration `yaml:"gracefull-duration" default:"10s"`
	}
)
