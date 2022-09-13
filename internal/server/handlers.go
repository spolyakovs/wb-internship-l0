package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/spolyakovs/wb-internship-l0/internal/store"
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
		case errors.Is(err, store.ErrSQLNotExist):
			srv.error(w, req, http.StatusBadRequest, store.ErrSQLNotExist)
			return
		case err != nil:
			srv.error(w, req, http.StatusInternalServerError, store.ErrSQLInternal)
			return
		}
		srv.respond(w, req, http.StatusOK, &order)
	}
}
