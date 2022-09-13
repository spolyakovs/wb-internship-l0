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
		pub.STANConnection.Close()
		os.Exit(0)
	}()

	pub.PublishRandomValid()
	pub.PublishRandomValid()
	pub.PublishRandomValid()
	pub.PublishRandomValid()

	pub.STANConnection.Close()
}
