package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spolyakovs/wb-internship-l0/src/store"
)

type server struct {
	router     *mux.Router
	store      store.Store
	sessionKey []byte
}

func newServer(st store.Store) *server {
	srv := &server{
		router: mux.NewRouter(),
		store:  st,
	}

	srv.configureRouter()

	return srv
}

func (server *server) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	server.router.ServeHTTP(writer, req)
}

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

func newDB(config Config) (*sqlx.DB, error) {
	dbURL := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DatabaseHost, config.DatabaseDBName, config.DatabaseUser, config.DatabasePassword, config.DatabaseSSLMode)
	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
