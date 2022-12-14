package store_test

import (
	"context"
	"testing"
)

func TestPaymentRepositoryCreate(t *testing.T) {
	ctx := context.Background()

	paymentRandom, err := getTestPayment()
	if err != nil {
		t.Errorf("couldn't create random payment: %v", err)
		return
	}

	if err := st.Payments().Create(ctx, &paymentRandom); err != nil {
		t.Errorf("didn't create payment: %v\n\tpayment: %+v", err, paymentRandom)
		return
	}
	if paymentRandom.ID == 0 {
		t.Errorf("created payment has no ID\n\tpayment: %+v", paymentRandom)
	}
}

func TestPaymentRepositoryFindById(t *testing.T) {
	ctx := context.Background()

	paymentRandom, err := getTestPayment()
	if err != nil {
		t.Errorf("couldn't create random payment: %v", err)
		return
	}

	if err := st.Payments().Create(ctx, &paymentRandom); err != nil {
		t.Errorf("couldn't create payment: %v\n\tpayment: %+v", err, paymentRandom)
		return
	}
	if paymentRandom.ID == 0 {
		t.Errorf("created payment has no ID\n\tpayment: %+v", paymentRandom)
		return
	}

	paymentFound, err := st.Payments().FindByID(ctx, paymentRandom.ID)
	if err != nil {
		t.Errorf("couldn't find payment: %v\n\tpayment: %+v", err, paymentRandom)
		return
	}

	if *paymentFound != paymentRandom {
		t.Errorf("found another payment\n\tfound: %+v\n\twanted: %+v", paymentFound, paymentRandom)
	}
}
