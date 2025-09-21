package config

import "os"

type Config struct {
	Addr string
}

func NewConfig() *Config {
	addr := os.Getenv("CONFIG_ADDR")
	if addr == "" {
		port := os.Getenv("CONFIG_PORT")
		if port == "" {
			port = "8080"
		}
		addr = ":" + port
	}
	return &Config{addr}
}
