package server_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/spolyakovs/wb-internship-l0/internal/app/publisher"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
	"github.com/spolyakovs/wb-internship-l0/internal/app/store"
)

var (
	srv *server.Server
	pub *publisher.Publisher
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	configPath := "../../../configs/local_test.toml"
	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		fmt.Printf("couldn't read config from file:%v\n\tpath:%s", err, configPath)
		return
	}

	db, err := server.NewDB(ctx, config)
	if err != nil {
		fmt.Printf("%v:\n\t%v", server.ErrDBCreate, err)
		return
	}
	defer db.Close()

	st := *store.New(db)

	srv = server.NewServer(ctx, st, config.LogLevel)

	go srv.StanSubscribe(ctx, config)

	pub, err = publisher.NewPublisher(config)
	if err != nil {
		fmt.Printf("coudn't initialize publisher:\n\t%v", err)
		return
	}

	defer pub.STANConnection.Close()

	os.Exit(m.Run())
}
