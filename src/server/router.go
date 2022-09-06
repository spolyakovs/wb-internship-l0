package server

func (srv *server) configureRouter() {
	srv.router.HandleFunc("/get", srv.handleGet()).Methods("GET")
}
