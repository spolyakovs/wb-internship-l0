package store_test

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

var st store.Store

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

	// if err := truncateTables(db); err != nil {
	// 	fmt.Printf("couldn't truncate db: %v", err)
	// 	return
	// }

	st = *store.New(db)

	if err := publisher.CustomFakerGenerator(); err != nil {
		fmt.Printf("faker generator err: %v", err)
		return
	}

	os.Exit(m.Run())
}

// func truncateTables(db *sql.DB) error {
// 	truncateQuery := `TRUNCATE TABLE order_items RESTART IDENTITY CASCADE;
// 	TRUNCATE TABLE orders RESTART IDENTITY CASCADE;
// 	TRUNCATE TABLE payments  RESTART IDENTITY CASCADE;
// 	TRUNCATE TABLE items  RESTART IDENTITY CASCADE;
// 	TRUNCATE TABLE deliveries  RESTART IDENTITY CASCADE;`
// 	_, err := db.Exec(truncateQuery)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
