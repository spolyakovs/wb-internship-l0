package server

import (
	"context"
	"encoding/json"
	"log"

	stan "github.com/nats-io/stan.go"
	"github.com/spolyakovs/wb-internship-l0/internal/model"
)

func stanSubscribe(ctx context.Context, config Config) error {
	sc, err := stan.Connect(config.STANClusterID, config.STANClientID)
	if err != nil {
		return err
	}
	defer sc.Close()

	errChan := make(chan error, 1)
	dataChan := make(chan model.Order, 10)

	go handleSTANErrors(errChan)
	go handleSTANdata(dataChan)

	sub, err := sc.Subscribe(config.STANChannel, func(m *stan.Msg) {
		// TODO: create DB entry and add into cache
		order := model.Order{}
		if err := json.Unmarshal(m.Data, &order); err != nil {
			errChan <- err
		} else {
			dataChan <- order
		}
	}, stan.StartWithLastReceived(), stan.DurableName(config.STANClientDurable))
	if err != nil {
		return err
	}

	<-ctx.Done()
	sub.Unsubscribe()
	close(errChan)
	close(dataChan)

	return nil
}

func handleSTANErrors(errChan <-chan error) {
	err := <-errChan
	log.Println(err.Error())
}

func handleSTANdata(dataChan <-chan model.Order) {
	order := <-dataChan
	log.Printf("%+v\n", order)
}
