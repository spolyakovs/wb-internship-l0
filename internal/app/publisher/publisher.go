package publisher

import (
	"encoding/json"
	"fmt"

	"github.com/bxcodec/faker/v4"
	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
)

type Publisher struct {
	logger         logrus.Logger
	STANChannel    string
	STANConnection stan.Conn
}

func NewPublisher(config server.Config) (*Publisher, error) {
	if err := CustomFakerGenerator(); err != nil {
		return nil, fmt.Errorf("faker generator err: %v", err)
	}

	sc, err := stan.Connect(config.STANClusterID, config.STANClientID)
	if err != nil {
		return nil, fmt.Errorf("%w: connection error: %v", server.ErrSTANInternal, err)
	}

	publisher := Publisher{
		logger:         *logrus.New(),
		STANConnection: sc,
		STANChannel:    config.STANChannel,
	}

	publisher.logger.Level = config.LogLevel

	return &publisher, nil
}

func (publisher Publisher) PublishRandomValid() (string, error) {
	fakeOrder := model.Order{}
	if err := faker.FakeData(&fakeOrder); err != nil {
		return "", fmt.Errorf("%w: %v", ErrFakerFakeData, err)
	}
	fakeOrderBytes, err := json.Marshal(fakeOrder)
	if err != nil {
		return "", fmt.Errorf("%w: %v\n\t", server.ErrJSONMarshal, err)
	}
	publisher.logger.Info("Publishing order with id:", fakeOrder.OrderUID)
	publisher.STANConnection.Publish(publisher.STANChannel, fakeOrderBytes)
	return fakeOrder.OrderUID, nil
}

func (publisher Publisher) PublishInvalid() error {
	fakeDelivery := model.Delivery{}
	if err := faker.FakeData(&fakeDelivery); err != nil {
		return fmt.Errorf("%w: %v", ErrFakerFakeData, err)
	}
	fakeDeliveryBytes, err := json.Marshal(fakeDelivery)
	if err != nil {
		return fmt.Errorf("%w: %v\n\t", server.ErrJSONMarshal, err)
	}
	publisher.logger.Info("Publishing delivery (invalid data)")
	publisher.STANConnection.Publish(publisher.STANChannel, fakeDeliveryBytes)
	return nil
}
