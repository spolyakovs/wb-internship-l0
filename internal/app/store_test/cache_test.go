package store_test

import (
	"context"
	"errors"
	"testing"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
	"github.com/spolyakovs/wb-internship-l0/internal/app/store"
)

func TestCache(t *testing.T) {
	ctx := context.Background()

	orderEmpty := model.Order{}
	orderRandom, err := getTestOrder()
	if err != nil {
		t.Errorf("couldn't create random order: %v", err)
		return
	}

	st.Cache().Put(&orderEmpty)

	orderEmptyFound, err := st.Cache().Get(ctx, orderEmpty.OrderUID)
	if err == nil {
		t.Errorf("got empty order from cache\n\torderFound: %+v", orderEmptyFound)
	} else {
		if !errors.Is(err, store.ErrNotExist) {
			t.Errorf("coudn't get empty order from cache: %v\n\t", err)
		}
	}

	orderFoundNotInDB, err := st.Cache().Get(ctx, orderRandom.OrderUID)
	if err == nil {
		t.Errorf("got order that isn't in DB from cache\n\torderFound: %+v", orderFoundNotInDB)
	} else {
		if !errors.Is(err, store.ErrNotExist) {
			t.Errorf("coudn't get order that isn't in DB from cache: %v\n\torder_uid: %+v", err, orderRandom.OrderUID)
		}
	}

	if err := st.Orders().Create(ctx, &orderRandom); err != nil {
		t.Errorf("didn't create order: %v\n\torder: %+v", err, orderRandom)
		return
	}

	orderFound1, err := st.Cache().Get(ctx, orderRandom.OrderUID)
	if err != nil {
		t.Errorf("coudn't get order that from cache: %v\n\torder_uid: %+v", err, orderRandom.OrderUID)
		return
	}
	if !equalOrders(*orderFound1, orderRandom) {
		t.Errorf("got another order from cache \n\twanted: %+v\n\tfound: %+v", orderRandom, orderFound1)
	}

	orderFound2, err := st.Cache().Get(ctx, orderRandom.OrderUID)
	if err != nil {
		t.Errorf("coudn't get order that from cache second time: %v\n\torder_uid: %+v", err, orderRandom.OrderUID)
		return
	}
	if !equalOrders(*orderFound2, orderRandom) {
		t.Errorf("got another order from cache second time\n\twanted: %+v\n\tfound: %+v", orderRandom, orderFound2)
	}
}
