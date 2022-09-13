package store_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/spolyakovs/wb-internship-l0/internal/app/publisher"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
	"github.com/spolyakovs/wb-internship-l0/internal/app/store"
)

var st store.Store

func TestMain(m *testing.M) {
	configPath := "../../../configs/local_test.toml"
	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := server.NewDB(ctx, config)
	if err != nil {
		fmt.Printf("%v:\n\t%v", server.ErrDBCreate, err)
		return
	}
	defer db.Close()

	st = *store.New(db)

	if err := publisher.CustomFakerGenerator(); err != nil {
		fmt.Printf("faker generator err: %v", err)
		return
	}

	os.Exit(m.Run())
}
