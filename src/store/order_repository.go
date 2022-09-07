package store

import (
	"context"
	"fmt"

	"github.com/spolyakovs/wb-internship-l0/src/model"
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
	createQuery := "INSERT INTO payments " +
		"(order_uid, track_number, entry, delivery_id, payment_id, " +
		"locale, internal_signature, customer_id, delivery_service, " +
		"shardkey, sm_id, date_created, oof_shard) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);"

	createRes, err := o.store.db.ExecContext(ctx, createQuery,
		order.OrderUid,
		order.TrackNumber,
		order.Entry,
		order.Delivery.Id,
		order.Payment.Id,
		order.Locale,
		order.InternalSignature,
		order.CustomerId,
		order.DeliveryService,
		order.ShardKey,
		order.SmId,
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
		connectQuery := "INSERT INTO order_items (order_uid, item_id) " +
			"VALUES ($1, $2);"
		connectRes, err := o.store.db.ExecContext(ctx, connectQuery,
			order.OrderUid,
			item.Id,
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

func (o orderRepository) FindById(ctx context.Context, id uint) (*model.Order, error) {
	order := model.Order{}
	var delivery_id, payment_id uint

	// Finding main order fields, delivery_id and payment_id
	findByIdQuery := "SELECT (order_uid, track_number, entry, delivery_id, payment_id, " +
		"locale, internal_signature, customer_id, delivery_service, " +
		"shardkey, sm_id, date_created, oof_shard) " +
		" FROM items WHERE id = $1 LIMIT 1;"

	if err := o.store.db.QueryRowContext(ctx,
		findByIdQuery,
		id,
	).Scan(
		&order.OrderUid,
		&order.TrackNumber,
		&order.Entry,
		&delivery_id,
		&payment_id,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerId,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmId,
		&order.DateCreated,
		&order.OofShard,
	); err != nil {
		return nil, err
	}

	// Finding delivery by id
	delivery, err := o.store.Deliveries().FindById(ctx, delivery_id)
	if err != nil {
		return nil, err
	}
	order.Delivery = delivery

	// And payment
	payment, err := o.store.Payments().FindById(ctx, payment_id)
	if err != nil {
		return nil, err
	}
	order.Payment = payment

	// and all items for this order
	items, err := o.store.Items().FindAllByOrderUid(ctx, order.OrderUid)
	if err != nil {
		return nil, err
	}
	order.Items = items

	return &order, nil
}
