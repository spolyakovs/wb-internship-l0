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

	"github.com/sirupsen/logrus"
	"github.com/spolyakovs/wb-internship-l0/internal/app/store"
)

func Start(config Config) error {
	logger := logrus.New()
	logger.Level = config.LogLevel

	logger.Info("Starting server")

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-appSignal
		logger.Info("Shutting down server")
		close(appSignal)
		stop()
		os.Exit(0)
	}()

	logger.Info("Connecting to DB")
	db, err := NewDB(ctx, config)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDBCreate, err)
	}
	defer db.Close()
	logger.Info("Successfully connected to DB")

	st := *store.New(db)

	srv := newServer(ctx, st)
	srv.logger.Level = config.LogLevel

	go srv.stanSubscribe(ctx, config)
	logger.Info("Server started")

	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(ctx context.Context, config Config) (*sql.DB, error) {
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
