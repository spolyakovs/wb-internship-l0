package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/spolyakovs/wb-internship-l0/src/store"
)

func Start(config Config) error {
	db, err := newDB(config)
	if err != nil {
		return err
	}

	defer db.Close()

	st := *store.New(db)

	srv := newServer(st)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(config Config) (*sql.DB, error) {
	dbURL := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DatabaseHost, config.DatabaseDBName, config.DatabaseUser, config.DatabasePassword, config.DatabaseSSLMode)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
