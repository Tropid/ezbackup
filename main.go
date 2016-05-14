package main

import (
	"flag"
	"log"

	"github.com/tropid/ezbackup/backup"
	"github.com/tropid/ezbackup/config"
)

func main() {
	var configFile = flag.String("config", "config.json", "Path to the config file.")
	flag.Parse()

	config, err := config.Load(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, g := range config.Groups {
		err = backup.Robocopy(g)

		if err != nil {
			log.Fatal(err)
		}
	}
}
