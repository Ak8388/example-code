package config

import (
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DB_HOST   string
	DB_PORT   string
	DB_USER   string
	DB_PASS   string
	DB_NAME   string
	DB_DRIVER string
}

type ApiConfig struct {
	API_PORT string
}

type Config struct {
	DBConfig
	ApiConfig
}

func (c *Config) readConfig() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	c.DBConfig = DBConfig{
		DB_HOST:   os.Getenv("DB_HOST"),
		DB_PORT:   os.Getenv("DB_PORT"),
		DB_USER:   os.Getenv("DB_USER"),
		DB_PASS:   os.Getenv("DB_PASS"),
		DB_NAME:   os.Getenv("DB_NAME"),
		DB_DRIVER: os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		API_PORT: os.Getenv("API_PORT"),
	}

	if c.DB_HOST == "" || c.DB_USER == "" || c.DB_PASS == "" || c.DB_NAME == "" || c.DB_PORT == "" || c.DB_DRIVER == "" || c.API_PORT == "" {
		panic("all env required")
	}
}

func NewConfig() *Config {
	cfg := &Config{}

	cfg.readConfig()

	return cfg
}
