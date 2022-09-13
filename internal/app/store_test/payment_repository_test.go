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
	} else {
		if paymentRandom.ID == 0 {
			t.Errorf("didn't create payment\n\tpayment: %+v", paymentRandom)
		}
	}
}

func TestPaymentRepositoryFindById(t *testing.T) {
	ctx := context.Background()

	paymentRandom, err := getTestPayment()
	if err != nil {
		t.Errorf("couldn't create random payment: %v", err)
	}

	if err := st.Payments().Create(ctx, &paymentRandom); err != nil {
		t.Errorf("couldn't create payment: %v\n\tpayment: %+v", err, paymentRandom)
		return
	} else {
		if paymentRandom.ID == 0 {
			t.Errorf("didn't create payment\n\tpayment: %+v", paymentRandom)
			return
		}
	}

	paymentFound, err := st.Payments().FindByID(ctx, paymentRandom.ID)
	if err != nil {
		t.Errorf("couldn't find payment: %v\n\tpayment: %+v", err, paymentRandom)
	}

	if *paymentFound != paymentRandom {
		t.Errorf("found another payment\n\twanted: %+v\n\tfound: %+v", paymentRandom, paymentFound)
	}
}
