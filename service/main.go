package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
)

func main() {
	configPath := "./configs/local.toml"
	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	config.STANClientID += "-subscriber"

	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}
