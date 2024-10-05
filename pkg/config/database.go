package pkgConfig

import "fmt"

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
}

func (db *Database) GenerateConnectionString() string {
	// example -> postgres://admin:admin@localhost:5432/template
	// example -> oracle://SIMASVIRTUALUAT:simasvirtual@10.32.1.19:1521/ORCL
	// return fmt.Sprintf("%+v://%+v:%+v@%+v:%+v/%+v", db.Driver, db.Username, db.Password, db.Host, db.Port, db.Schema)
	return fmt.Sprintf("user=%+v password=%+v dbname=%+v sslmode=disable host=%+v port=%+v", db.Username, db.Password, db.Schema, db.Host, db.Port)
}
