package server

import "github.com/BurntSushi/toml"

type Config struct {
	BindAddr          string `toml:"BIND_ADDR"`
	LogLevel          string `toml:"LOG_LEVEL"`
	DatabaseHost      string `toml:"DATABASE_HOST"`
	DatabaseDBName    string `toml:"DATABASE_DB"`
	DatabaseUser      string `toml:"DATABASE_USER"`
	DatabasePassword  string `toml:"DATABASE_PASSWORD"`
	DatabaseSSLMode   string `toml:"DATABASE_SSLMODE"`
	STANAddr          string `toml:"STAN_ADDR"`
	STANClusterID     string `toml:"STAN_CLUSTER_ID"`
	STANClientID      string `toml:"STAN_CLIENT_ID"`
	STANClientDurable string `toml:"STAN_CLIENT_DURABLE"`
	STANChannel       string `toml:"STAN_CHANNEL"`
}

func MakeConfigFromFile(path string) (Config, error) {
	config := Config{}

	if _, err := toml.DecodeFile(path, &config); err != nil {
		return config, err
	}

	return config, nil
}

// func NewConfig() *Config {
// 	return &Config{
// 		BindAddr: ":8080",
// 		LogLevel: "debug",
// 		// DatabaseHost: "",
// 		// DatabaseDBName: "",
// 		// DatabaseUser: "",
// 		// DatabasePassword: "",
// 		// DatabaseSSLMode: "",
// 		STANAddr: ":4222",
// 		// STANClisterID: "",
// 		// STANClientID: "",
// 		// STANClientDurable: "",
// 		// STANChannel: "",
// 	}
// }
