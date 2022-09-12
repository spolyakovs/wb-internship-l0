package store

import (
	"context"
	"sync"

	"github.com/spolyakovs/wb-internship-l0/internal/model"
)

type cache struct {
	sync.RWMutex
	store *Store
	data  map[string]*model.Order
}

func (c *cache) Add(order *model.Order) {
	c.Lock()
	c.data[order.OrderUID] = order
	c.Unlock()
}

func (c *cache) Get(ctx context.Context, order_uid string) (*model.Order, error) {
	c.RLock()
	order, ok := c.data[order_uid]
	c.RUnlock()
	if ok {
		return order, nil
	}

	order, err := c.store.Orders().FindByID(ctx, order_uid)
	if err != nil {
		return nil, err
	}
	c.Add(order)

	return order, nil
}
