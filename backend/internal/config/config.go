package config

import "os"

type Config struct {
    DBDSN string
    Port  string
}

func Load() *Config {
    return &Config{
        DBDSN: getEnv("DB_DSN", "root:root@tcp^(127.0.0.1:3306^)/mall?charset=utf8mb4^&parseTime=True^&loc=Local"),
        Port:  getEnv("PORT", "8080"),
    }
}

func getEnv(k, def string) string {
    if v := os.Getenv(k); v = "" {
        return v
    }
    return def
}
