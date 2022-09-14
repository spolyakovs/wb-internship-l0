package store_test

import (
	"context"
	"testing"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
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
		return
	}

	if itemRandom.ID == 0 {
		t.Errorf("created item has no ID\n\titem: %+v", itemRandom)
	}
}

func TestItemRepositoryFindById(t *testing.T) {
	ctx := context.Background()

	itemRandom, err := getTestItem()
	if err != nil {
		t.Errorf("couldn't create random item: %v", err)
		return
	}

	if err := st.Items().Create(ctx, &itemRandom); err != nil {
		t.Errorf("couldn't create item: %v\n\titem: %+v", err, itemRandom)
		return
	}

	if itemRandom.ID == 0 {
		t.Errorf("created item has no ID\n\titem: %+v", itemRandom)
		return
	}

	itemFound, err := st.Items().FindByID(ctx, itemRandom.ID)
	if err != nil {
		t.Errorf("couldn't find item: %v\n\titem: %+v", err, itemRandom)
		return
	}

	if *itemFound != itemRandom {
		t.Errorf("found another item\n\tfound: %+v\n\twanted: %+v", itemFound, itemRandom)
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
		return
	}

	if err := st.Orders().Create(ctx, &orderRandom); err != nil {
		t.Errorf("couldn't create order: %v\n\titem: %+v", err, orderRandom)
		return
	}

	itemsFound, err := st.Items().FindAllByOrderUID(ctx, orderRandom.OrderUID)
	if err != nil {
		t.Errorf("couldn't find items by order_uid: %v\n\torder_uid: %+v", err, orderRandom.OrderUID)
		return
	}

	if !equalItems(itemsFound, orderRandom.Items) {
		t.Errorf("found other items for order\n\tfound: %+v\n\twanted: %+v", itemsFound, orderRandom.Items)
	}
}

// compare itemsFound and itemsWanted as sets, not slices (sequence of elements is irrelevant)
func equalItems(itemsFound, itemsWanted []*model.Item) bool {
	if len(itemsFound) != len(itemsWanted) {
		return false
	}

	diff := make(map[uint]int)

	for i := 0; i < len(itemsFound); i++ {
		// initiate map[key]value pair for IDs
		if _, ok := diff[itemsFound[i].ID]; !ok {
			diff[itemsFound[i].ID] = 0
		}
		if _, ok := diff[itemsWanted[i].ID]; !ok {
			diff[itemsFound[i].ID] = 0
		}
		// increase value if item in itemsFound, decrease if in orderRandom.Items
		// will become 0 if exists in both slices
		diff[itemsFound[i].ID] += 1
		diff[itemsWanted[i].ID] -= 1
		if diff[itemsWanted[i].ID] == 0 {
			delete(diff, itemsWanted[i].ID)
		}
	}

	// if slices are identical, all values will become 0 and will be deleted
	return len(diff) == 0
}
