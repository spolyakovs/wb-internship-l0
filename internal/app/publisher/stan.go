package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bxcodec/faker/v4"
	stan "github.com/nats-io/stan.go"
	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
)

func stanPublishValid(ctx context.Context, config server.Config) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second+10*time.Millisecond)
	defer cancel()
	sc, err := stan.Connect(config.STANClusterID, config.STANClientID)
	if err != nil {
		cancel()
		return fmt.Errorf("%w: %v", server.ErrSTANInternal, err)
	}
	defer sc.Close()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(time.Second):
			fakeOrder := model.Order{}
			if err := faker.FakeData(&fakeOrder); err != nil {
				return fmt.Errorf("%w: %v", ErrFakerFakeData, err)
			}
			fakeOrderBytes, err := json.Marshal(fakeOrder)
			if err != nil {
				return fmt.Errorf("%w: %v\n\t", server.ErrJSONMarshal, err)
			}
			fmt.Println("Publishing order with id:", fakeOrder.OrderUID)
			sc.Publish(config.STANChannel, fakeOrderBytes)
		}
	}
}
