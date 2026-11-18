package config

import "os"

type Config struct {
	HTTPAddr string
	MySQLDSN string
}

func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	return Config{
		HTTPAddr: addr,
		MySQLDSN: os.Getenv("MYSQL_DSN"),
	}
}
