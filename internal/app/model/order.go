package model

import "fmt"

type Order struct {
	OrderUID          string    `json:"order_uid" faker:"uuid_digit,unique"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          *Delivery `json:"delivery"`
	Payment           *Payment  `json:"payment"`
	Items             []*Item   `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       string    `json:"date_created" faker:"customFakerTimestamp"` // format RFC3339 without timezone (2006-01-02T15:04:05Z)
	OofShard          string    `json:"oof_shard"`
}

func (order *Order) Validate() error {
	if order.OrderUID == "" {
		err := fmt.Errorf("%w: OrderUID", ErrMissingRequiredField)
		return fmt.Errorf("Order %w: %v", ErrValidation, err)
	}
	if order.Delivery == nil {
		err := fmt.Errorf("%w: Delivery", ErrMissingRequiredField)
		return fmt.Errorf("Order %w: %v", ErrValidation, err)
	}
	if order.Payment == nil {
		err := fmt.Errorf("%w: Payment", ErrMissingRequiredField)
		return fmt.Errorf("Order %w: %v", ErrValidation, err)
	}
	return nil
}
