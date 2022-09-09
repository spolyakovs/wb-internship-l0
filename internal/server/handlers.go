package server

import (
	"context"
	"net/http"
)

func (srv *server) handleHello(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// delivery := model.Delivery{}
		// if err := srv.store.Deliveries().Create(ctx, &delivery); err != nil {
		// 	fmt.Printf("/nError:/n/t%w", err)
		// }
		srv.respond(w, req, http.StatusOK, "Hello")
		// return
	}
}
