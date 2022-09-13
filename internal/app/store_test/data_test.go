package store_test

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
	"github.com/spolyakovs/wb-internship-l0/internal/app/publisher"
)

func getTestDelivery() (model.Delivery, error) {
	fakeDelivery := model.Delivery{}
	if err := faker.FakeData(&fakeDelivery); err != nil {
		return fakeDelivery, fmt.Errorf("%w (delivery): %v", publisher.ErrFakerFakeData, err)
	}

	return fakeDelivery, nil
}

func getTestItem() (model.Item, error) {
	fakeItem := model.Item{}
	if err := faker.FakeData(&fakeItem); err != nil {
		return fakeItem, fmt.Errorf("%w (item): %v", publisher.ErrFakerFakeData, err)
	}

	return fakeItem, nil
}

func getTestOrder() (model.Order, error) {
	fakeOrder := model.Order{}
	if err := faker.FakeData(&fakeOrder); err != nil {
		return fakeOrder, fmt.Errorf("%w (order): %v", publisher.ErrFakerFakeData, err)
	}

	return fakeOrder, nil
}

func getTestPayment() (model.Payment, error) {
	fakePayment := model.Payment{}
	if err := faker.FakeData(&fakePayment); err != nil {
		return fakePayment, fmt.Errorf("%w (payment): %v", publisher.ErrFakerFakeData, err)
	}

	return fakePayment, nil
}
