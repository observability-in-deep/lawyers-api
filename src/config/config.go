package config

import (
	"os"
	"strconv"
)

type Config struct {
	Host          string
	Port          int
	ServiceName   string
	ListenAddress string
	IsLocal       bool
	OtlpEndpoint  string
}

func NewConfig() *Config {
	return &Config{
		ServiceName:   "lawyers-api",
		ListenAddress: getEnv("GO_LISTEN_ADDRESS", ":3001"),
		IsLocal:       getEnvAsBool("IS_LOCAL", true),
		OtlpEndpoint:  getEnv("OTLP_ENDPOINT", "localhost:4317"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultValue
}
