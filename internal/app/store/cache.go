package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
)

type cache struct {
	sync.RWMutex
	store *Store
	data  map[string]*model.Order
}

func (c *cache) Put(order *model.Order) {
	// put only valid orders
	if err := order.Validate(); err != nil {
		return
	}

	c.Lock()
	c.data[order.OrderUID] = order
	c.Unlock()
}

func (c *cache) Get(ctx context.Context, order_uid string) (*model.Order, error) {
	// if order alreay in cache, then return it
	c.RLock()
	order, ok := c.data[order_uid]
	c.RUnlock()
	if ok {
		return order, nil
	}

	// else get order from DB, put it in cache and return
	order, err := c.store.Orders().FindByID(ctx, order_uid)
	if err != nil {
		return nil, fmt.Errorf("couldn't get order from cache with id: %v\n\t%w", order_uid, err)
	}
	c.Put(order)

	return order, nil
}
