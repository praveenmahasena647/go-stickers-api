package dbs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	postgres *sql.DB
)

func Connect() error {
	var err error
	var dns = fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	postgres, err = sql.Open("postgres", dns)
	if err != nil {
		return err
	}
	return postgres.Ping()
}
