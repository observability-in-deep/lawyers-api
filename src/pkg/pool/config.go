package pool

import (
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConfig() *Config {
	return &Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "lawyer"),
		Password: getEnv("DB_PASSWORD", "lawyer"),
		DBName:   getEnv("DB_NAME", "lawyer"),
	}
}

func (c *Config) CreateString() string {
	return "host=" + c.Host + " port=" + string(c.Port) + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
