package store

import (
	"context"

	"github.com/spolyakovs/wb-internship-l0/src/model"
)

type deliveryRepository struct {
	store *Store
}

func (d deliveryRepository) Create(ctx context.Context, delivery *model.Delivery) error {
	createQuery := "INSERT INTO deliveries " +
		"(name, phone, zip, city, address, region, email) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;"

	if err := d.store.db.QueryRowContext(ctx, createQuery,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	).Scan(&delivery.Id); err != nil {
		return err
	}

	return nil
}

func (d deliveryRepository) FindById(ctx context.Context, id uint) (*model.Delivery, error) {
	delivery := model.Delivery{}
	findByIdQuery := "SELECT (id, name, phone, zip, city, address, region, email) " +
		"FROM deliveries WHERE id = $1 LIMIT 1;"

	if err := d.store.db.QueryRowContext(ctx, findByIdQuery,
		id,
	).Scan(
		&delivery.Id,
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	); err != nil {
		return nil, err
	}

	return &delivery, nil
}
