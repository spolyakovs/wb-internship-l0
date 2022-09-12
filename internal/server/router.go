package server

import "context"

func (srv *server) configureRouter(ctx context.Context) {
	srv.router.HandleFunc("/", srv.handleGetOrder(ctx)).Methods("GET")
}
