package store

import (
	"context"
	"fmt"

	"github.com/spolyakovs/wb-internship-l0/internal/app/model"
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
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	).Scan(&payment.ID); err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return fmt.Errorf("couldn't create payment: %+v\n\t%w", payment, err)
	}

	return nil
}

func (p paymentRepository) FindByID(ctx context.Context, id uint) (*model.Payment, error) {
	payment := model.Payment{}
	findByIdQuery := "SELECT " +
		"id, transaction, request_id, currency, provider, amount, " +
		"payment_dt, bank, delivery_cost, goods_total, custom_fee " +
		"FROM payments WHERE id = $1 LIMIT 1;"

	if err := p.store.db.QueryRowContext(ctx,
		findByIdQuery,
		id,
	).Scan(
		&payment.ID,
		&payment.Transaction,
		&payment.RequestID,
		&payment.Currency,
		&payment.Provider,
		&payment.Amount,
		&payment.PaymentDt,
		&payment.Bank,
		&payment.DeliveryCost,
		&payment.GoodsTotal,
		&payment.CustomFee,
	); err != nil {
		err = fmt.Errorf("%w: %v", ErrSQLInternal, err)
		return nil, fmt.Errorf("couldn't find payment with id: %v\n\t%w", id, err)
	}

	return &payment, nil
}
