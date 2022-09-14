package server_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/spolyakovs/wb-internship-l0/internal/app/publisher"
	"github.com/spolyakovs/wb-internship-l0/internal/app/server"
)

var (
	baseURL string
	pub     *publisher.Publisher
)

func TestMain(m *testing.M) {
	configPath := "../../../configs/local_test.toml"
	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		fmt.Printf("couldn't read config from file:%v\n\tpath:%s", err, configPath)
		return
	}

	if err := server.Start(config); err != nil {
		fmt.Printf("coudn't start server:\n\t%v", err)
		return
	}

	baseURL = config.BindAddr

	pub, err = publisher.NewPublisher(config)
	if err != nil {
		fmt.Printf("coudn't initialize publisher:\n\t%v", err)
		return
	}

	defer pub.STANConnection.Close()

	os.Exit(m.Run())
}
