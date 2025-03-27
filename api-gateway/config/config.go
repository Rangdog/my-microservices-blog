package config

import "github.com/joho/godotenv"
type Config struct{
	ConsulAddr string
}

func LoadConfig() (*Config, error){
	godotenv.Load()
	return &Config{
		ConsulAddr: "consul:8500",
	}, nil
}