package pool

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "lawyer-user",
		Password: "lawyer-password",
		DBName:   "lawyers",
	}
}

func (c *Config) CreateString() string {
	return "host=" + c.Host + " port=" + string(c.Port) + " user=" + c.User + " password=" + c.Password + " dbname=" + c.DBName
}
