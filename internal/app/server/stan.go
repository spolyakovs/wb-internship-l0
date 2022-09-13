package server

import (
	"context"
	"encoding/json"
	"fmt"

	stan "github.com/nats-io/stan.go"
	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
)

func (srv server) stanSubscribe(ctx context.Context, config Config) {
	errChan := make(chan error, 1)
	defer close(errChan)

	dataChan := make(chan model.Order, 10)
	defer close(dataChan)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go srv.handleSTANErrors(ctx, errChan)
	go srv.handleSTANdata(ctx, dataChan, errChan)

	srv.logger.Infof("Connecting to STAN with clusterID: %v, clientID: %v", config.STANClusterID, config.STANClientID)

	sc, err := stan.Connect(config.STANClusterID, config.STANClientID)
	if err != nil {
		errChan <- fmt.Errorf("%w: connection error: %v", ErrSTANInternal, err)
		return
	}
	defer sc.Close()

	srv.logger.Infof("Subscribing to STAN channel: %v with clientDurable:%v", config.STANChannel, config.STANClientDurable)
	sub, err := sc.Subscribe(config.STANChannel, func(m *stan.Msg) {
		order := model.Order{}
		if err := json.Unmarshal(m.Data, &order); err != nil {
			errChan <- fmt.Errorf("%w: %v: %v", ErrSTANReceived, ErrJSONUnmarshal, err)
			return
		}

		if err := order.Validate(); err != nil {
			errChan <- fmt.Errorf("%w: %v", ErrSTANReceived, err)
			return
		}

		srv.logger.Infof("Received order: %+v\n", order)
		dataChan <- order
	}, stan.StartWithLastReceived(), stan.DurableName(config.STANClientDurable))
	if err != nil {
		errChan <- fmt.Errorf("%w: subscribe error: %v", ErrSTANInternal, err)
		return
	}
	srv.logger.Info("Successfully subscribed to STAN")

	<-ctx.Done()
	sub.Unsubscribe()
}

func (srv server) handleSTANErrors(ctx context.Context, errChan <-chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		case err := <-errChan:
			if err != nil { // needed because can receive nils when channel is closed
				srv.logger.Warnln("STAN error:", err)
			}
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
				errChan <- fmt.Errorf("%w: %v", ErrSTANReceived, err)
			}
		}
	}
}
