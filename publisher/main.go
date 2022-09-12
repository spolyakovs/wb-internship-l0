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
		log.Fatalf("Config err:%v", err)
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
		log.Fatalf("Faker generator err:%v", err)
	}

	if err := stanPublishRandom(ctx, config); err != nil {
		log.Fatalf("STAN publish err:%v", err)
	}
}
