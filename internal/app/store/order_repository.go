package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
)

type orderRepository struct {
	store *Store
}

func (o orderRepository) Create(ctx context.Context, order *model.Order) error {
	// Creating sub-entities
	if err := o.store.Deliveries().Create(ctx, order.Delivery); err != nil {
		return fmt.Errorf("couldn't create order.delivery\n\t%w", err)
	}
	if err := o.store.Payments().Create(ctx, order.Payment); err != nil {
		return fmt.Errorf("couldn't create order.payment\n\t%w", err)
	}
	for _, item := range order.Items {
		if err := o.store.Items().Create(ctx, item); err != nil {
			return fmt.Errorf("couldn't create order.items\n\t%w", err)
		}
	}

	// Creating main entity
	createQuery := "INSERT INTO orders " +
		"(order_uid, track_number, entry, delivery_id, payment_id, " +
		"locale, internal_signature, customer_id, delivery_service, " +
		"shardkey, sm_id, date_created, oof_shard) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);"

	createRes, err := o.store.db.ExecContext(ctx, createQuery,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Delivery.ID,
		order.Payment.ID,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.ShardKey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't create order: %+v\n\t%w", order, err)
	}
	rows, err := createRes.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSQLInternal, err)
	}
	if rows != 1 {
		err = fmt.Errorf("expected to affect 1 row, affected %d", rows)
		return fmt.Errorf("%w: %v", ErrSQLInternal, err)
	}

	// Connecting order with its items through "order_items" entity
	for _, item := range order.Items {
		connectQuery := "INSERT INTO order_items " +
			"(order_uid, item_id) " +
			"VALUES ($1, $2);"
		connectRes, err := o.store.db.ExecContext(ctx, connectQuery,
			order.OrderUID,
			item.ID,
		)
		if err != nil {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
			return fmt.Errorf("couldn't create order_items entities\n\t%w", err)
		}
		rows, err := connectRes.RowsAffected()
		if err != nil {
			return fmt.Errorf("%w: %v", ErrSQLInternal, err)
		}
		if rows != 1 {
			err = fmt.Errorf("expected to affect 1 row, affected %d", rows)
			return fmt.Errorf("%w: %v", ErrSQLInternal, err)
		}
	}

	return nil
}

func (o orderRepository) FindByID(ctx context.Context, order_uid string) (*model.Order, error) {
	order := model.Order{}
	var delivery_id, payment_id uint

	// Finding main order fields, delivery_id and payment_id
	findByIdQuery := "SELECT " +
		"order_uid, track_number, entry, delivery_id, payment_id, " +
		"locale, internal_signature, customer_id, delivery_service, " +
		"shardkey, sm_id, date_created, oof_shard " +
		"FROM orders WHERE order_uid = $1 LIMIT 1;"

	if err := o.store.db.QueryRowContext(ctx,
		findByIdQuery,
		order_uid,
	).Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&delivery_id,
		&payment_id,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard,
	); err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		} else {
			err = fmt.Errorf("%w: %v", ErrSQLNotExist, err)
		}
		return nil, fmt.Errorf("couldn't find order with order_uid: %v\n\t%w", order_uid, err)
	}

	// Finding delivery by id
	delivery, err := o.store.Deliveries().FindByID(ctx, delivery_id)
	if err != nil {
		return nil, fmt.Errorf("couldn't find order.delivery with order_uid: %v\n\t%w", order_uid, err)
	}
	order.Delivery = delivery

	// And payment
	payment, err := o.store.Payments().FindByID(ctx, payment_id)
	if err != nil {
		return nil, fmt.Errorf("couldn't find order.payment with order_uid: %v\n\t%w", order_uid, err)
	}
	order.Payment = payment

	// and all items for this order
	items, err := o.store.Items().FindAllByOrderUID(ctx, order.OrderUID)
	if err != nil {
		return nil, fmt.Errorf("couldn't find order.items with order_uid: %v\n\t%w", order_uid, err)
	}
	order.Items = items

	return &order, nil
}
