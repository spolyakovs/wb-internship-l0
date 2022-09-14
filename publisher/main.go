package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spolyakovs/wb-internship-l0/internal/app/publisher"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
)

func main() {
	configPath := "./configs/local.toml"
	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		log.Fatalf("config err: %v", err)
	}
	config.STANClientID += "-publisher"

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt, syscall.SIGTERM)

	pub, err := publisher.NewPublisher(config)
	if err != nil {
		log.Fatalf("creating publisher error: %v", err)
	}

	go func() {
		<-appSignal
		os.Exit(0)
	}()

	for i := 0; i < 4; i++ {
		if _, err := pub.PublishRandomValid(); err != nil {
			pub.STANConnection.Close()
			log.Fatalf("error while publishing random order: %v", err)
		}
	}

	pub.STANConnection.Close()
}
