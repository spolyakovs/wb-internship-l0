package store_test

import (
	"context"
	"testing"
)

func TestDeliveryRepositoryCreate(t *testing.T) {
	ctx := context.Background()

	deliveryRandom, err := getTestDelivery()
	if err != nil {
		t.Errorf("couldn't create random delivery: %v", err)
		return
	}

	if err := st.Deliveries().Create(ctx, &deliveryRandom); err != nil {
		t.Errorf("didn't create delivery: %v\n\tdelivery: %+v", err, deliveryRandom)
		return
	}
	if deliveryRandom.ID == 0 {
		t.Errorf("created delivery has no ID\n\tdelivery: %+v", deliveryRandom)
	}
}

func TestDeliveryRepositoryFindById(t *testing.T) {
	ctx := context.Background()

	deliveryRandom, err := getTestDelivery()
	if err != nil {
		t.Errorf("couldn't create random delivery: %v", err)
		return
	}

	if err := st.Deliveries().Create(ctx, &deliveryRandom); err != nil {
		t.Errorf("couldn't create delivery: %v\n\tdelivery: %+v", err, deliveryRandom)
		return
	}
	if deliveryRandom.ID == 0 {
		t.Errorf("created delivery has no ID\n\tdelivery: %+v", deliveryRandom)
		return
	}

	deliveryFound, err := st.Deliveries().FindByID(ctx, deliveryRandom.ID)

	if err != nil {
		t.Errorf("couldn't find delivery: %v\n\tdelivery: %+v", err, deliveryRandom)
		return
	}

	if *deliveryFound != deliveryRandom {
		t.Errorf("found another delivery\n\twanted: %+v\n\tfound: %+v", deliveryRandom, deliveryFound)
	}
}
