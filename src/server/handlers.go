package server

import (
	"net/http"
)

func (srv *server) handleGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		return
	}
}
