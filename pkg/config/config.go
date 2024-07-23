package pkgConfig

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		AppConfig AppConfig  `yaml:"app"`
		Database  Database   `yaml:"database"`
		Auth      AuthConfig `yaml:"auth"`
		Redis     Redis      `yaml:"redis"`
	}
)

func NewConfig() (Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
		os.Setenv("ENV", "development")
	}

	filePath := fmt.Sprintf("./variables/%s.yaml", env)

	file, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}
