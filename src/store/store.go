package store

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
	// userRepository *UserRepository
}

func New(db *sql.DB) *Store {
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
