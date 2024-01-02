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
	postgres, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return postgres.Ping()
}
