package main

import (
	"log"

	"github.com/praveenmahasena647/go-stickers/cmd/app"
)

func main() {
	var appErr = app.Start()
	if appErr != nil {
		log.Fatal(appErr)
	}
}
