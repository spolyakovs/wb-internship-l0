package store

import (
	"context"

	"github.com/spolyakovs/wb-internship-l0/src/model"
)

type itemRepository struct {
	store *Store
}

func (i itemRepository) Create(ctx context.Context, item *model.Item) error {
	createQuery := "INSERT INTO items " +
		"(chrt_id, track_number, price, rid, name, sale, nm_id, brand, status) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;"

	if err := i.store.db.QueryRowContext(ctx, createQuery,
		item.ChrtId,
		item.TrackNumber,
		item.Price,
		item.RId,
		item.Name,
		item.Sale,
		item.NmId,
		item.Brand,
		item.Status,
	).Scan(&item.Id); err != nil {
		return err
	}

	return nil
}

func (i itemRepository) FindById(ctx context.Context, id uint) (*model.Item, error) {
	item := model.Item{}
	findByIdQuery := "SELECT (id, chrt_id, track_number, price, rid, name, sale, nm_id, brand, status)" +
		" FROM items WHERE id = $1 LIMIT 1;"

	if err := i.store.db.QueryRowContext(ctx,
		findByIdQuery,
		id,
	).Scan(
		&item.Id,
		&item.ChrtId,
		&item.TrackNumber,
		&item.Price,
		&item.RId,
		&item.Name,
		&item.Sale,
		&item.NmId,
		&item.Brand,
		&item.Status,
	); err != nil {
		return nil, err
	}

	return &item, nil
}

func (i itemRepository) FindAllByOrderUid(ctx context.Context, orderUid string) ([]*model.Item, error) {
	items := make([]*model.Item, 0)
	findByIdQuery := "SELECT (items.id, items.chrt_id, items.track_number, items.price, " +
		"items.rid, items.name, items.sale, items.nm_id, items.brand, items.status) " +
		"FROM items LEFT JOIN order_items ON items.id = order_items.item_id " +
		"WHERE order_items.order_uid = $1;"

	rows, err := i.store.db.QueryContext(ctx,
		findByIdQuery,
		orderUid,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		item := model.Item{}
		if err := rows.Scan(
			&item.Id,
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.RId,
			&item.Name,
			&item.Sale,
			&item.NmId,
			&item.Brand,
			&item.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
