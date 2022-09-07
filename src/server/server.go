package server

import (
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spolyakovs/wb-internship-l0/src/store"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(st store.Store) *server {
	srv := &server{
		router: mux.NewRouter(),
		store:  st,
	}

	srv.configureRouter()

	return srv
}

func (server *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

func (server *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	server.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (server *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
