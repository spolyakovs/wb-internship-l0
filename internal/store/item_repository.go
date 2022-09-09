package store

import (
	"context"

	"github.com/spolyakovs/wb-internship-l0/internal/model"
)

type itemRepository struct {
	store *Store
}

func (i itemRepository) Create(ctx context.Context, item *model.Item) error {
	createQuery := "INSERT INTO items " +
		"(chrt_id, track_number, price, rid, name, sale, size, " +
		"total_price, nm_id, brand, status) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;"

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
		return err
	}

	return nil
}

func (i itemRepository) FindByID(ctx context.Context, id uint) (*model.Item, error) {
	item := model.Item{}
	findByIdQuery := "SELECT " +
		"(id, chrt_id, track_number, price, rid, name, sale, size, " +
		"total_price, nm_id, brand, status) " +
		" FROM items WHERE id = $1 LIMIT 1;"

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
		return nil, err
	}

	return &item, nil
}

func (i itemRepository) FindAllByOrderUID(ctx context.Context, orderUid string) ([]*model.Item, error) {
	items := make([]*model.Item, 0)
	findByIdQuery := "SELECT " +
		"(items.id, items.chrt_id, items.track_number, items.price, " +
		"items.rid, items.name, items.sale, items.size, items.total_price" +
		"items.nm_id, items.brand, items.status) " +
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
			&item.ID,
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.NmID,
			&item.Brand,
			&item.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
