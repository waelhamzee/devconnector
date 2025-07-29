package config

import "os"

type Config struct {
	Port string
	Env  string
}

func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("GIN_MODE", "debug"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
