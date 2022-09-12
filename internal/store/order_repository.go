package store

import (
	"context"
	"fmt"

	"github.com/spolyakovs/wb-internship-l0/internal/model"
)

type orderRepository struct {
	store *Store
}

func (o orderRepository) Create(ctx context.Context, order *model.Order) error {
	// Creating sub-entities
	if err := o.store.Deliveries().Create(ctx, order.Delivery); err != nil {
		return err
	}
	if err := o.store.Payments().Create(ctx, order.Payment); err != nil {
		return err
	}
	for _, item := range order.Items {
		if err := o.store.Items().Create(ctx, item); err != nil {
			return err
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
		return err
	}
	rows, err := createRes.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
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
			return err
		}
		rows, err := connectRes.RowsAffected()
		if err != nil {
			return err
		}
		if rows != 1 {
			return fmt.Errorf("expected to affect 1 row, affected %d", rows)
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
		return nil, err
	}

	// Finding delivery by id
	delivery, err := o.store.Deliveries().FindByID(ctx, delivery_id)
	if err != nil {
		return nil, err
	}
	order.Delivery = delivery

	// And payment
	payment, err := o.store.Payments().FindByID(ctx, payment_id)
	if err != nil {
		return nil, err
	}
	order.Payment = payment

	// and all items for this order
	items, err := o.store.Items().FindAllByOrderUID(ctx, order.OrderUID)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return &order, nil
}
