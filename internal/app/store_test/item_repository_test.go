package store_test

import (
	"context"
	"testing"
)

func TestItemRepositoryCreate(t *testing.T) {
	ctx := context.Background()

	itemRandom, err := getTestItem()
	if err != nil {
		t.Errorf("couldn't create random item: %v", err)
		return
	}

	if err := st.Items().Create(ctx, &itemRandom); err != nil {
		t.Errorf("didn't create item: %v\n\titem: %+v", err, itemRandom)
	} else {
		if itemRandom.ID == 0 {
			t.Errorf("didn't create item\n\titem: %+v", itemRandom)
		}
	}
}

func TestItemRepositoryFindById(t *testing.T) {
	ctx := context.Background()

	itemRandom, err := getTestItem()
	if err != nil {
		t.Errorf("couldn't create random item: %v", err)
	}

	if err := st.Items().Create(ctx, &itemRandom); err != nil {
		t.Errorf("couldn't create item: %v\n\titem: %+v", err, itemRandom)
	} else {
		if itemRandom.ID == 0 {
			t.Errorf("didn't create item\n\titem: %+v", itemRandom)
		}
	}

	itemFound, err := st.Items().FindByID(ctx, itemRandom.ID)
	if err != nil {
		t.Errorf("couldn't find item: %v\n\titem: %+v", err, itemRandom)
	}

	if *itemFound != itemRandom {
		t.Errorf("found another item\n\twanted: %+v\n\tfound: %+v", itemRandom, itemFound)
	}
}

func TestItemRepositoryFindAllByOrderUID(t *testing.T) {
	ctx := context.Background()

	itemsFoundNotExist, err := st.Items().FindAllByOrderUID(ctx, "not_exist")
	if err != nil {
		t.Errorf("couldn't find items by non-existent order_uid: %v\n\t", err)
	} else {
		if len(itemsFoundNotExist) != 0 {
			t.Errorf("found items by non-existent order_uid: %v\n\titems: %+v", err, itemsFoundNotExist)
		}
	}

	orderRandom, err := getTestOrder()
	if err != nil {
		t.Errorf("couldn't create random order: %v", err)
	}

	if err := st.Orders().Create(ctx, &orderRandom); err != nil {
		t.Errorf("couldn't create order: %v\n\titem: %+v", err, orderRandom)
		return
	}

	itemsFound, err := st.Items().FindAllByOrderUID(ctx, orderRandom.OrderUID)
	if err != nil {
		t.Errorf("couldn't find items by order_uid: %v\n\torder_uid: %+v", err, orderRandom.OrderUID)
	}

	// TODO: think about checking items if they are in different sequence
	for i := range itemsFound {
		if *itemsFound[i] != *orderRandom.Items[i] {
			t.Errorf("found another item\n\twanted: %+v\n\tfound: %+v", orderRandom.Items[i], itemsFound[i])
		}
	}
}
