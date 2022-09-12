package server

import (
	"context"
	"errors"
	"net/http"
)

func (srv *server) handleGetOrder(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		orderUID := req.URL.Query().Get("id")
		if orderUID == "" {
			srv.error(w, req, http.StatusBadRequest, errors.New("no order id in request"))
			return
		}

		order, err := srv.store.Cache().Get(ctx, orderUID)
		if err != nil {
			srv.error(w, req, http.StatusInternalServerError, err)
			return
		}
		srv.respond(w, req, http.StatusOK, &order)
		// return
	}
}
