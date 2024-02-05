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
	return fmt.Sprintf("user=%+v password=%+v dbname=%+v sslmode=disable host=%+v port=%+v", db.Username, db.Password, db.Schema, db.Host, db.Port)
}
