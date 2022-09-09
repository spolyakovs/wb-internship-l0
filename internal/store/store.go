package store

import (
	"database/sql"
)

type Store struct {
	db                 *sql.DB
	deliveryRepository *deliveryRepository
	itemRepository     *itemRepository
	paymentRepository  *paymentRepository
	orderRepository    *orderRepository
}

func New(db *sql.DB) *Store {
	newStore := &Store{
		db: db,
	}

	return newStore
}

func (st *Store) Deliveries() deliveryRepository {
	if st.deliveryRepository == nil {
		st.deliveryRepository = &deliveryRepository{
			store: st,
		}
	}

	return *(st.deliveryRepository)
}

func (st *Store) Items() itemRepository {
	if st.itemRepository == nil {
		st.itemRepository = &itemRepository{
			store: st,
		}
	}

	return *(st.itemRepository)
}

func (st *Store) Payments() paymentRepository {
	if st.paymentRepository == nil {
		st.paymentRepository = &paymentRepository{
			store: st,
		}
	}

	return *(st.paymentRepository)
}

func (st *Store) Orders() orderRepository {
	if st.orderRepository == nil {
		st.orderRepository = &orderRepository{
			store: st,
		}
	}

	return *(st.orderRepository)
}
