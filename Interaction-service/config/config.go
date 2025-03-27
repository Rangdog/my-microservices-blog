package config

import "os"

type Config struct {
	MySQLDSN string
	Port     string
	JWTSecret string
}

func LoadConfig() *Config {
	return &Config{
		MySQLDSN: os.Getenv("MYSQL_DSN"),
		Port: os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}