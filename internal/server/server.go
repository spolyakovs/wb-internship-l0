package server

import (
	"context"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spolyakovs/wb-internship-l0/internal/store"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(ctx context.Context, st store.Store) *server {
	srv := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  st,
	}

	srv.configureRouter(ctx)

	return srv
}

func (server *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	server.router.ServeHTTP(w, req)
}

func (server *server) error(w http.ResponseWriter, req *http.Request, code int, err error) {
	server.respond(w, req, code, map[string]string{"error": err.Error()})
}

func (server *server) respond(w http.ResponseWriter, req *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
