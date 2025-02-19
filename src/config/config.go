package config

type Config struct {
	Host          string
	Port          int
	ServiceName   string
	ListenAddress string
	IsLocal       bool
}

func NewConfig() *Config {
	return &Config{
		Host:          "localhost",
		Port:          3000,
		ServiceName:   "lawyers-api",
		ListenAddress: "localhost:3001",
		IsLocal:       true,
	}
}
