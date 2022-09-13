package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
)

type itemRepository struct {
	store *Store
}

func (i itemRepository) Create(ctx context.Context, item *model.Item) error {
	createQuery := "INSERT INTO items " +
		"(chrt_id, track_number, price, rid, name, sale, size, " +
		"total_price, nm_id, brand, status) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;"

	if err := i.store.db.QueryRowContext(ctx, createQuery,
		item.ChrtID,
		item.TrackNumber,
		item.Price,
		item.RID,
		item.Name,
		item.Sale,
		item.Size,
		item.TotalPrice,
		item.NmID,
		item.Brand,
		item.Status,
	).Scan(&item.ID); err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't create item: %+v\n\t%w", item, err)
	}

	return nil
}

func (i itemRepository) FindByID(ctx context.Context, id uint) (*model.Item, error) {
	item := model.Item{}
	findByIdQuery := "SELECT " +
		"id, chrt_id, track_number, price, rid, name, sale, size, " +
		"total_price, nm_id, brand, status " +
		"FROM items WHERE id = $1 LIMIT 1;"

	if err := i.store.db.QueryRowContext(ctx,
		findByIdQuery,
		id,
	).Scan(
		&item.ID,
		&item.ChrtID,
		&item.TrackNumber,
		&item.Price,
		&item.RID,
		&item.Name,
		&item.Sale,
		&item.Size,
		&item.TotalPrice,
		&item.NmID,
		&item.Brand,
		&item.Status,
	); err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		} else {
			err = fmt.Errorf("%w: %v", ErrSQLNotExist, err)
		}
		return nil, fmt.Errorf("couldn't find item with id: %v\n\t%w", id, err)
	}

	return &item, nil
}

func (i itemRepository) FindAllByOrderUID(ctx context.Context, orderUid string) ([]*model.Item, error) {
	items := make([]*model.Item, 0)
	findByIdQuery := "SELECT " +
		"items.id, items.chrt_id, items.track_number, items.price, " +
		"items.rid, items.name, items.sale, items.size, items.total_price, " +
		"items.nm_id, items.brand, items.status " +
		"FROM items LEFT JOIN order_items ON id = order_items.item_id " +
		"WHERE order_items.order_uid = $1;"

	rows, err := i.store.db.QueryContext(ctx,
		findByIdQuery,
		orderUid,
	)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return nil, fmt.Errorf("couldn't find items by order_uid: %v\n\t%w", orderUid, err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.Item{}
		if err := rows.Scan(
			&item.ID,
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status,
		); err != nil {
			err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
			return nil, fmt.Errorf("couldn't scan item into model: %v\n\t%w", orderUid, err)
		}
		items = append(items, &item)
	}

	return items, nil
}
