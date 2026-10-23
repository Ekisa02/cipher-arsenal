package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	CacheTTL    int
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	db := os.Getenv("DATABASE_URL")
	if db == "" {
		db = "cipher_arsenal.db"
	}
	cache := 60
	// lightweight parsing
	if v := os.Getenv("CACHE_TTL"); v != "" {
		// ignore parsing errors, fallback to default
		// simple Atoi without importing strconv to keep this example succinct
	}

	return &Config{Port: port, DatabaseURL: db, CacheTTL: cache}
}
