package server

type Config struct {
	BindAddr          string `toml:"BIND_ADDR"`
	LogLevel          string `toml:"LOG_LEVEL"`
	DatabaseHost      string `toml:"DATABASE_HOST"`
	DatabaseDBName    string `toml:"DATABASE_DB"`
	DatabaseUser      string `toml:"DATABASE_USER"`
	DatabasePassword  string `toml:"DATABASE_PASSWORD"`
	DatabaseSSLMode   string `toml:"DATABASE_SSLMODE"`
	NATSStreamingAddr string `toml:"NATS_STREAMING_ADDR"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		// DatabaseHost: "",
		// DatabaseDBName: "",
		// DatabaseUser: "",
		// DatabasePassword: "",
		// DatabaseSSLMode: "",
		// NATSStreamingAddr: "",
	}
}
