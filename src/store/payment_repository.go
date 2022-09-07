package store

import (
	"context"

	"github.com/spolyakovs/wb-internship-l0/src/model"
)

type paymentRepository struct {
	store *Store
}

func (p paymentRepository) Create(ctx context.Context, payment *model.Payment) error {
	createQuery := "INSERT INTO payments " +
		"(transaction, request_id, currency, provider, amount, " +
		"payment_dt, bank, delivery_cost, goods_total, custom_fee) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;"

	if err := p.store.db.QueryRowContext(ctx, createQuery,
		payment.Transaction,
		payment.RequestId,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	).Scan(&payment.Id); err != nil {
		return err
	}

	return nil
}

func (p paymentRepository) FindById(ctx context.Context, id uint) (*model.Payment, error) {
	payment := model.Payment{}
	findByIdQuery := "SELECT (id, chrt_id, track_number, price, rid, name, sale, nm_id, brand, status)" +
		" FROM items WHERE id = $1 LIMIT 1;"

	if err := p.store.db.QueryRowContext(ctx,
		findByIdQuery,
		id,
	).Scan(
		&payment.Id,
		&payment.Transaction,
		&payment.RequestId,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	); err != nil {
		return nil, err
	}

	return &payment, nil
}
