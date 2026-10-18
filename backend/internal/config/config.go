package config

import "os"

type Config struct {
	DBDSN string
	Port  string
}

func Load() *Config {
	return &Config{
		DBDSN: getEnv("DB_DSN", "host=localhost user=postgres password=root dbname=mall port=5432 sslmode=disable"),
		Port:  getEnv("PORT", "8080"),
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
