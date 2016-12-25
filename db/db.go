package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

// SetupDatabaseConnection opens up & returns a SQL connection
func SetupDatabaseConnection() (*sql.DB, error) {
	var err error
	var dsn = fmt.Sprintf(
		"%s:%s@%s/%s",
		os.Getenv("ANA_DATABASE_USER"),
		os.Getenv("ANA_DATABASE_PASSWORD"),
		os.Getenv("ANA_DATABASE_HOST"),
		os.Getenv("ANA_DATABASE_NAME"),
	)

	Conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Open doesn't open a connection right away. Validate DSN by calling Ping().
	err = Conn.Ping()
	if err != nil {
		return nil, err
	}

	return Conn, nil
}
