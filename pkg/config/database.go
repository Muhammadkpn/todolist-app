package pkgConfig

import "fmt"

type DatabaseCfg struct {
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

type Database struct {
	Template DatabaseCfg `yaml:"template"`
	Oracle   DatabaseCfg `yaml:"oracle"`
}

func (db *DatabaseCfg) GenerateConnectionString() string {
	// example -> postgres://admin:admin@localhost:5432/template
	// example -> oracle://SIMASVIRTUALUAT:simasvirtual@10.32.1.19:1521/ORCL
	return fmt.Sprintf("%+v://%+v:%+v@%+v:%+v/%+v", db.Driver, db.Username, db.Password, db.Host, db.Port, db.Schema)
}
