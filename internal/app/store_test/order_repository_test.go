package store_test

import (
	"context"
	"errors"
	"testing"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
	"github.com/spolyakovs/wb-internship-l0/internal/app/store"
)

func TestOrderRepositoryCreate(t *testing.T) {
	ctx := context.Background()

	orderRandom, err := getTestOrder()
	if err != nil {
		t.Errorf("couldn't create random order: %v", err)
		return
	}

	if err := st.Orders().Create(ctx, &orderRandom); err != nil {
		t.Errorf("didn't create order: %v\n\torder: %+v", err, orderRandom)
		return
	}

	if err := st.Orders().Create(ctx, &orderRandom); err == nil {
		t.Errorf("created order second time\n\torder: %+v", orderRandom)
	} else {
		if !errors.Is(err, store.ErrAlreadyExist) {
			t.Errorf("coudn't create order second time: %v\n\torder: %+v", err, orderRandom)
		}
	}
}

func TestOrderRepositoryFindById(t *testing.T) {
	ctx := context.Background()

	orderRandom, err := getTestOrder()
	if err != nil {
		t.Errorf("couldn't create random order: %v", err)
		return
	}

	if err := st.Orders().Create(ctx, &orderRandom); err != nil {
		t.Errorf("couldn't create order: %v\n\torder: %+v", err, orderRandom)
		return
	}

	orderFound, err := st.Orders().FindByID(ctx, orderRandom.OrderUID)
	if err != nil {
		t.Errorf("couldn't find order: %v\n\torder: %+v", err, orderRandom)
		return
	}

	if !equalOrders(*orderFound, orderRandom) {
		t.Errorf("found another order\n\twanted: %+v\n\tfound: %+v", orderRandom, orderFound)
	}
}

// needed because of deliveries, payments and items
// order contains pointers, but need to compare actual models
func equalOrders(orderFound, orderWanted model.Order) bool {
	if orderFound.OrderUID != orderWanted.OrderUID {
		return false
	}
	if orderFound.TrackNumber != orderWanted.TrackNumber {
		return false
	}
	if orderFound.Entry != orderWanted.Entry {
		return false
	}
	if *orderFound.Delivery != *orderWanted.Delivery {
		return false
	}
	if *orderFound.Payment != *orderWanted.Payment {
		return false
	}
	if !equalItems(orderFound.Items, orderWanted.Items) {
		return false
	}
	if orderFound.Locale != orderWanted.Locale {
		return false
	}
	if orderFound.InternalSignature != orderWanted.InternalSignature {
		return false
	}
	if orderFound.CustomerID != orderWanted.CustomerID {
		return false
	}
	if orderFound.DeliveryService != orderWanted.DeliveryService {
		return false
	}
	if orderFound.ShardKey != orderWanted.ShardKey {
		return false
	}
	if orderFound.SmID != orderWanted.SmID {
		return false
	}
	if orderFound.DateCreated != orderWanted.DateCreated {
		return false
	}
	if orderFound.OofShard != orderWanted.OofShard {
		return false
	}
	return true
}
