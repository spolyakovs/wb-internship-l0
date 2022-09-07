package server

import (
	"net/http"
)

func (srv *server) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv.respond(w, r, http.StatusOK, "Hello")
		return
	}
}
