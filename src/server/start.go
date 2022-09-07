package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spolyakovs/wb-internship-l0/src/store"
)

// TODO: wrap and refactor errors
// TODO: connect to nets-streaming

func Start(config Config) error {

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-appSignal
		stop()
		os.Exit(0)
	}()

	db, err := newDB(ctx, config)
	if err != nil {
		return err
	}

	defer db.Close()

	st := *store.New(db)

	srv := newServer(ctx, st)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(ctx context.Context, config Config) (*sql.DB, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	dbURL := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=%s",
		config.DatabaseHost, config.DatabaseDBName, config.DatabaseUser, config.DatabasePassword, config.DatabaseSSLMode)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
