package main

import (
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
	"github.com/spolyakovs/wb-internship-l0/src/server"
)

func main() {
	configPath := "./configs/local.toml"
	config := server.NewConfig()

	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	if err := server.Start(*config); err != nil {
		log.Fatal(err)
	}
}
