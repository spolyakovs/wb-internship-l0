package store

import "errors"

var (
	ErrNotExist     = errors.New("couldn't find this record")
	ErrAlreadyExist = errors.New("record with this ID already exist")
	ErrSQLInternal  = errors.New("internal sql error")
)
