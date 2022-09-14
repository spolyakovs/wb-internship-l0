package server

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func (srv *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := srv.logger.WithFields(logrus.Fields{
			"remote_addr": req.RemoteAddr,
		})

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, req)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"%s %s completed with %d %s in %v",
			req.Method,
			req.RequestURI,
			rw.code,
			http.StatusText(rw.code),
			time.Since(start),
		)
	})
}
