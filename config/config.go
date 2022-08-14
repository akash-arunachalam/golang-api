package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "34.93.109.172",
			Port:     3306,
			Username: "akash",
			Password: "Akash#99",
			Name:     "ak",
			Charset:  "utf8",
		},
	}
}
