package server

func (srv *server) configureRouter() {
	srv.router.HandleFunc("/", srv.handleHello()).Methods("GET")
}
