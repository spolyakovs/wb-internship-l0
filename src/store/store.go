package store

import (
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
	// userRepository *UserRepository
}

func New(db *sqlx.DB) *Store {
	newStore := &Store{
		db: db,
	}

	return newStore
}

// func (st *Store) Users() store.UserRepository {
// 	if st.userRepository != nil {
// 		return st.userRepository
// 	}

// 	st.userRepository = &UserRepository{
// 		store: st,
// 	}

// 	return st.userRepository
// }
