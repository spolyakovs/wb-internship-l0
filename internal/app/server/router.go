package server

import "context"

func (srv *Server) configureRouter(ctx context.Context) {
	srv.router.Use(srv.logRequest)
	srv.router.HandleFunc("/", srv.handleGetOrder(ctx)).Methods("GET")
}
