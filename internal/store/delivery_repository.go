package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spolyakovs/wb-internship-l0/internal/model"
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
	).Scan(&delivery.ID); err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't create delivery: %+v\n\t%w", delivery, err)
	}

	return nil
}

func (d deliveryRepository) FindByID(ctx context.Context, id uint) (*model.Delivery, error) {
	delivery := model.Delivery{}
	findByIdQuery := "SELECT " +
		"id, name, phone, zip, city, address, region, email " +
		"FROM deliveries WHERE id = $1 LIMIT 1;"

	if err := d.store.db.QueryRowContext(ctx, findByIdQuery,
		id,
	).Scan(
		&delivery.ID,
		&delivery.Name,
		&delivery.Phone,
		&delivery.Zip,
		&delivery.City,
		&delivery.Address,
		&delivery.Region,
		&delivery.Email,
	); err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		} else {
			err = fmt.Errorf("%w: %v", ErrSQLNotExist, err)
		}
		return nil, fmt.Errorf("couldn't find delivery with id: %v\n\t%w", id, err)
	}

	return &delivery, nil
}
