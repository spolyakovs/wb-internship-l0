package store

import "errors"

var (
	ErrSQLNotExist = errors.New("couldn't find this record")
	ErrSQLInternal = errors.New("internal sql error")
)
