package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	stan "github.com/nats-io/stan.go"
	"github.com/spolyakovs/wb-internship-l0/internal/model"
)

func (srv server) stanSubscribe(ctx context.Context, config Config) {
	errChan := make(chan error, 1)
	dataChan := make(chan model.Order, 10)

	go srv.handleSTANErrors(ctx, errChan)
	go srv.handleSTANdata(ctx, dataChan, errChan)

	sc, err := stan.Connect(config.STANClusterID, config.STANClientID)
	if err != nil {
		errChan <- err
		return
	}
	defer sc.Close()
	defer close(dataChan)
	defer close(errChan)

	sub, err := sc.Subscribe(config.STANChannel, func(m *stan.Msg) {
		order := model.Order{}
		if err := json.Unmarshal(m.Data, &order); err != nil {
			errChan <- err
		} else {
			fmt.Println("Received order")
			dataChan <- order
		}
	}, stan.StartWithLastReceived(), stan.DurableName(config.STANClientDurable))
	if err != nil {
		errChan <- err
		return
	}

	<-ctx.Done()
	sub.Unsubscribe()
}

func (srv server) handleSTANErrors(ctx context.Context, errChan <-chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-errChan:
			log.Println(err.Error())
		}
	}
}

func (srv server) handleSTANdata(
	ctx context.Context,
	dataChan <-chan model.Order,
	errChan chan<- error,
) {
	for {
		select {
		case <-ctx.Done():
			return
		case order := <-dataChan:
			if err := srv.store.Orders().Create(ctx, &order); err != nil {
				errChan <- err
			}
		}
	}
}
