package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spolyakovs/wb-internship-l0/internal/server"
)

func main() {
	configPath := "./configs/local.toml"
	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	config.STANClientID += "-publisher"

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-appSignal
		stop()
		os.Exit(0)
	}()

	if err := customFakerGenerator(); err != nil {
		log.Fatal(err)
	}

	if err := stanPublishRandom(ctx, config); err != nil {
		log.Fatal(err)
	}
}
