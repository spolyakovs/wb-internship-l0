package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/spolyakovs/wb-internship-l0/internal/app/store"
)

func (srv *server) handleGetOrder(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		orderUID := req.URL.Query().Get("id")
		if orderUID == "" {
			srv.error(w, req, http.StatusBadRequest, errors.New("no order id in request"))
			return
		}

		order, err := srv.store.Cache().Get(ctx, orderUID)
		switch {
		case errors.Is(err, store.ErrNotExist):
			srv.error(w, req, http.StatusBadRequest, store.ErrNotExist)
			return
		case err != nil:
			srv.error(w, req, http.StatusInternalServerError, store.ErrSQLInternal)
			return
		}
		srv.respond(w, req, http.StatusOK, &order)
	}
}
