package dbs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	postgres *sql.DB
)

func Connect() error {
	var err error
	var dns = fmt.Sprintf("user=postgres password=CJ dbname=stickers host=localhost port=5432 sslmode=disable")
	postgres, err = sql.Open("postgres", dns)
	if err != nil {
		return err
	}
	return postgres.Ping()
}
