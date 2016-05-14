package main

import (
	"log"

	"github.com/tropid/ezbackup/backup"
	"github.com/tropid/ezbackup/config"
)

func main() {
	config, err := config.Load("config.json")

	for _, g := range config.Groups {
		err = backup.Robocopy(g)

		if err != nil {
			log.Fatal(err)
		}
	}
}
