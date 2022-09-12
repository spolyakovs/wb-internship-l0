package store

import (
	"database/sql"

	"github.com/spolyakovs/wb-internship-l0/internal/model"
)

type Store struct {
	db                 *sql.DB
	deliveryRepository *deliveryRepository
	itemRepository     *itemRepository
	paymentRepository  *paymentRepository
	orderRepository    *orderRepository
	cache              *cache
}

func New(db *sql.DB) *Store {
	newStore := &Store{
		db: db,
	}

	return newStore
}

func (st *Store) Deliveries() *deliveryRepository {
	if st.deliveryRepository == nil {
		st.deliveryRepository = &deliveryRepository{
			store: st,
		}
	}

	return st.deliveryRepository
}

func (st *Store) Items() *itemRepository {
	if st.itemRepository == nil {
		st.itemRepository = &itemRepository{
			store: st,
		}
	}

	return st.itemRepository
}

func (st *Store) Payments() *paymentRepository {
	if st.paymentRepository == nil {
		st.paymentRepository = &paymentRepository{
			store: st,
		}
	}

	return st.paymentRepository
}

func (st *Store) Orders() *orderRepository {
	if st.orderRepository == nil {
		st.orderRepository = &orderRepository{
			store: st,
		}
	}

	return st.orderRepository
}

func (st *Store) Cache() *cache {
	if st.cache == nil {
		st.cache = &cache{
			store: st,
			data:  make(map[string]*model.Order),
		}
	}

	return st.cache
}
