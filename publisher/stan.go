package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v4"
	stan "github.com/nats-io/stan.go"
	"github.com/spolyakovs/wb-internship-l0/internal/model"
	"github.com/spolyakovs/wb-internship-l0/internal/server"
)

func stanPublishRandom(ctx context.Context, config server.Config) error {
	// TODO: a lot of returns and canceling ???
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	sc, err := stan.Connect(config.STANClusterID, config.STANClientID)
	if err != nil {
		cancel()
		return err
	}
	defer sc.Close()

	for {
		select {
		case <-ctx.Done():
			cancel()
			return nil
		case <-time.After(time.Second):
			fakeOrder := model.Order{}
			if err := faker.FakeData(&fakeOrder); err != nil {
				cancel()
				return err
			}
			fakeOrderBytes, err := json.Marshal(fakeOrder)
			if err != nil {
				cancel()
				return err
			}
			fmt.Println("Publishing order")
			sc.Publish(config.STANChannel, fakeOrderBytes)
		}
	}
}
