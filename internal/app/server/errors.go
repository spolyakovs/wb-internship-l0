package server

import "errors"

var (
	ErrSTANInternal = errors.New("internal STAN error")
	ErrSTANReceived = errors.New("wrong data format received via STAN, expected model.Order")

	ErrJSONUnmarshal = errors.New("couldn't unmarshal model into JSON")
	ErrJSONMarshal   = errors.New("couldn't marshal JSON into model")
	ErrDBCreate      = errors.New("couldn't create DB connection")
)
