package pkgConfig

import "fmt"

type RedisCfg struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Main RedisCfg `yaml:"main"`
}

func (redis *RedisCfg) GenerateConnectionString() string {
	// example -> redis://user:password@localhost:6379
	// example -> redis://localhost:6379
	if redis.Username == "" && redis.Password == "" {
		return fmt.Sprintf("redis://%+v:%+v", redis.Host, redis.Port)
	}

	return fmt.Sprintf("redis://%+v:%+v@%+v:%+v", redis.Username, redis.Password, redis.Host, redis.Port)
}
