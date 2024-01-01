package app

import (
	"github.com/praveenmahasena647/go-stickers/cmd/app/api"
	"github.com/praveenmahasena647/go-stickers/cmd/app/dbs"
)

func Start() error {
	var postgresConnection = dbs.Connect()
	if postgresConnection != nil {
		return postgresConnection
	}
	var r = api.New()
	if err := r.Run(":42069"); err != nil {
		return err
	}

	return nil
}
