package config

import "os"

type Config struct {
	DBDSN  string
	Port   string
	Secret string // 新增
}

func Load() *Config {
	return &Config{
		DBDSN:  getEnv("DB_DSN", "..."),
		Port:   getEnv("PORT", "8080"),
		Secret: getEnv("JWT_SECRET", "mall-freshco-super-secret"), // 新增默认密钥
	}
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
