package injection

import (
	pkgConfig "base/pkg/config"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitSqlx(cfg pkgConfig.Config) (*sqlx.DB, error) {
	connectionString := cfg.Database.GenerateConnectionString()
	fmt.Println(connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Configure *sql.DB options
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Ping the database to check if the connection is valid
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	// Wrap *sql.DB with sqlx.NewDb to get a *sqlx.DB
	dbx := sqlx.NewDb(db, "postgres")

	return dbx, nil
}
